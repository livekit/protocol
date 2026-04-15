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

	result := subprocessResult{
		EffectiveCPU: EffectiveCPURequest(),
		NumCPU:       c.NumCPU(),
		PlatformCPU:  plat.numCPU(),
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
