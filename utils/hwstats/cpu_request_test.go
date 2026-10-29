package hwstats

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/frostbyte73/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Mock platform monitor
// ---------------------------------------------------------------------------

type mockPlatform struct {
	nCPU float64
}

func (m *mockPlatform) getCPUIdle() (float64, error) { return m.nCPU, nil }
func (m *mockPlatform) numCPU() float64              { return m.nCPU }

func newTestCPUStats(plat platformCPUMonitor) *CPUStats {
	return &CPUStats{
		platform:        plat,
		warningThrottle: core.NewThrottle(time.Minute),
		closeChan:       make(chan struct{}),
	}
}

// ---------------------------------------------------------------------------
// NumCPU without override (no env var set in this process)
// ---------------------------------------------------------------------------

func TestNumCPU_NoOverride(t *testing.T) {
	c := newTestCPUStats(&mockPlatform{nCPU: 64})
	assert.Equal(t, 64.0, c.NumCPU())

	c2 := newTestCPUStats(&mockPlatform{nCPU: 8.5})
	assert.Equal(t, 8.5, c2.NumCPU())
}

// ---------------------------------------------------------------------------
// Subprocess tests for EffectiveCPURequest and NumCPU with override
//
// sync.Once means we can only test one env-var value per process.
// Each test case spawns a subprocess with the desired env.
// ---------------------------------------------------------------------------

const subprocessEnvKey = "HWSTATS_TEST_SUBPROCESS"

type subprocessResult struct {
	EffectiveCPU float64 `json:"effective_cpu"`
	NumCPU       float64 `json:"num_cpu"`
	PlatformCPU  float64 `json:"platform_cpu"`
	HostNumCPU   float64 `json:"host_num_cpu"`

	Samples map[string]sampleResult `json:"samples"`
}

type sampleResult struct {
	Idle         float64 `json:"idle"`
	CallbackIdle float64 `json:"callback_idle"`
	Load         float64 `json:"load"`
}

// sampleOnce takes one sample from plat and returns the stored idle, the idle passed
// to the idle callback, and the load.
func sampleOnce(plat platformCPUMonitor, hostCPU bool) sampleResult {
	var r sampleResult
	c := newTestCPUStats(plat)
	c.hostCPU = hostCPU
	c.idleCallback = func(idle float64) { r.CallbackIdle = idle }
	c.sampleCPUIdle()
	r.Idle, r.Load = c.GetCPUIdle(), c.GetCPULoad()
	return r
}

func TestEffectiveCPURequest(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		expectedCPU float64
	}{
		{"valid integer", "4", 4},
		{"valid float", "3.33", 3.33},
		{"valid large", "15", 15},
		{"valid fractional", "0.5", 0.5},
		{"empty", "", 0},
		{"negative", "-1", 0},
		{"zero", "0", 0},
		{"invalid string", "abc", 0},
		{"invalid mixed", "4abc", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := runSubprocess(t, tt.envValue)
			assert.InDelta(t, tt.expectedCPU, result.EffectiveCPU, 0.001)
		})
	}
}

func TestNumCPU_HostIgnoresCPURequest(t *testing.T) {
	result := runSubprocess(t, "2")
	assert.Equal(t, 2.0, result.NumCPU)
	assert.Equal(t, result.PlatformCPU, result.HostNumCPU)
}

// With LIVEKIT_CPU_REQUEST=8, idle sampled against a larger platform count (the cgroup
// quota, or every CPU without a limit) must be expressed against the request, or
// GetCPULoad (idle / request) reads 0 on a busy pod.
func TestSampleCPUIdle_CPURequest(t *testing.T) {
	result := runSubprocess(t, "8")

	for _, tc := range []struct {
		name string
		idle float64
		load float64
	}{
		// request below the platform count: idle is request minus used
		{"no limit, 4 of 64 busy", 4, 0.5},
		{"no limit, 10 of 64 busy", 0, 1},
		{"limit 16, 4 busy", 4, 0.5},

		// request equal to the platform count: nothing to rescale
		{"limit 8, 4 busy", 4, 0.5},

		// request above the platform count (Downward API rounds 7500m up to 8):
		// idle is left as measured, so a saturated cgroup still reads as full load
		{"limit 7.5, idle", 7.5, 0.0625},
		{"limit 7.5, 2 busy", 5.5, 0.3125},
		{"limit 7.5, saturated", 0, 1},

		// WithHostCPU ignores the request
		{"host cpu, 4 of 64 busy", 60, 0.0625},
	} {
		s, ok := result.Samples[tc.name]
		require.True(t, ok, tc.name)
		assert.InDelta(t, tc.idle, s.Idle, 0.001, tc.name)
		assert.InDelta(t, tc.idle, s.CallbackIdle, 0.001, tc.name)
		assert.InDelta(t, tc.load, s.Load, 0.001, tc.name)
	}
}

func runSubprocess(t *testing.T, cpuRequestEnv string) subprocessResult {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=TestSubprocessWorker", "-test.v")

	// Build clean env
	var env []string
	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, "LIVEKIT_CPU_REQUEST=") &&
			!strings.HasPrefix(e, subprocessEnvKey+"=") {
			env = append(env, e)
		}
	}
	env = append(env, subprocessEnvKey+"=1")
	if cpuRequestEnv != "" {
		env = append(env, "LIVEKIT_CPU_REQUEST="+cpuRequestEnv)
	}
	cmd.Env = env

	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Fatalf("subprocess failed: %s\nstderr: %s", err, exitErr.Stderr)
		}
		t.Fatalf("subprocess failed: %s", err)
	}

	var result subprocessResult
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "{") {
			require.NoError(t, json.Unmarshal([]byte(line), &result))
			return result
		}
	}
	t.Fatalf("no JSON result in subprocess output: %s", string(out))
	return result
}

// TestSubprocessWorker runs inside spawned subprocesses only.
func TestSubprocessWorker(t *testing.T) {
	if os.Getenv(subprocessEnvKey) != "1" {
		t.Skip("only runs as subprocess")
	}

	plat := &mockPlatform{nCPU: float64(runtime.NumCPU())}
	c := newTestCPUStats(plat)

	host := newTestCPUStats(plat)
	host.hostCPU = true

	result := subprocessResult{
		EffectiveCPU: EffectiveCPURequest(),
		NumCPU:       c.NumCPU(),
		PlatformCPU:  plat.numCPU(),
		HostNumCPU:   host.NumCPU(),
		Samples: map[string]sampleResult{
			"no limit, 4 of 64 busy":  sampleOnce(&fakeCPUMonitor{n: 64, idle: 60}, false),
			"no limit, 10 of 64 busy": sampleOnce(&fakeCPUMonitor{n: 64, idle: 54}, false),
			"limit 16, 4 busy":        sampleOnce(&fakeCPUMonitor{n: 16, idle: 12}, false),
			"limit 8, 4 busy":         sampleOnce(&fakeCPUMonitor{n: 8, idle: 4}, false),
			"limit 7.5, idle":         sampleOnce(&fakeCPUMonitor{n: 7.5, idle: 7.5}, false),
			"limit 7.5, 2 busy":       sampleOnce(&fakeCPUMonitor{n: 7.5, idle: 5.5}, false),
			"limit 7.5, saturated":    sampleOnce(&fakeCPUMonitor{n: 7.5, idle: 0}, false),
			"host cpu, 4 of 64 busy":  sampleOnce(&fakeCPUMonitor{n: 64, idle: 60}, true),
		},
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
