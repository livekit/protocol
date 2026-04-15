package maxprocs_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const subprocessEnvKey = "MAXPROCS_TEST_SUBPROCESS"

type subprocessResult struct {
	GOMAXPROCS int `json:"gomaxprocs"`
}

func TestInstall(t *testing.T) {
	hostCPU := runtime.NumCPU()

	tests := []struct {
		name             string
		cpuRequest       string // LIVEKIT_CPU_REQUEST value ("" = unset)
		gomaxprocsEnv    string // GOMAXPROCS value ("" = unset)
		expectedMaxProcs int
	}{
		{
			name:             "no env vars set",
			expectedMaxProcs: hostCPU,
		},
		{
			name:             "request caps GOMAXPROCS down",
			cpuRequest:       "4",
			expectedMaxProcs: 4,
		},
		{
			name:             "fractional request rounds up",
			cpuRequest:       "3.33",
			expectedMaxProcs: 4,
		},
		{
			name:             "request of 1",
			cpuRequest:       "1",
			expectedMaxProcs: 1,
		},
		{
			name:             "request larger than host CPUs (no change)",
			cpuRequest:       "9999",
			expectedMaxProcs: hostCPU,
		},
		{
			name:             "explicit GOMAXPROCS lower than request (respected)",
			cpuRequest:       "8",
			gomaxprocsEnv:    "4",
			expectedMaxProcs: 4,
		},
		{
			name:             "explicit GOMAXPROCS higher than request (capped)",
			cpuRequest:       "4",
			gomaxprocsEnv:    "16",
			expectedMaxProcs: 4,
		},
		{
			name:             "explicit GOMAXPROCS equal to request",
			cpuRequest:       "8",
			gomaxprocsEnv:    "8",
			expectedMaxProcs: 8,
		},
		{
			name:             "explicit GOMAXPROCS without request (unchanged)",
			gomaxprocsEnv:    "4",
			expectedMaxProcs: 4,
		},
		{
			name:             "invalid request ignored",
			cpuRequest:       "abc",
			expectedMaxProcs: hostCPU,
		},
		{
			name:             "negative request ignored",
			cpuRequest:       "-2",
			expectedMaxProcs: hostCPU,
		},
		{
			name:             "zero request ignored",
			cpuRequest:       "0",
			expectedMaxProcs: hostCPU,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := runSubprocess(t, tt.cpuRequest, tt.gomaxprocsEnv)
			assert.Equal(t, tt.expectedMaxProcs, result.GOMAXPROCS)
		})
	}
}

func runSubprocess(t *testing.T, cpuRequest, gomaxprocs string) subprocessResult {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=TestSubprocessWorker", "-test.v")

	// Build clean env: start from current, filter out the vars we control
	var env []string
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "LIVEKIT_CPU_REQUEST=") ||
			strings.HasPrefix(e, "GOMAXPROCS=") ||
			strings.HasPrefix(e, subprocessEnvKey+"=") {
			continue
		}
		env = append(env, e)
	}
	env = append(env, subprocessEnvKey+"=1")
	if cpuRequest != "" {
		env = append(env, "LIVEKIT_CPU_REQUEST="+cpuRequest)
	}
	if gomaxprocs != "" {
		env = append(env, "GOMAXPROCS="+gomaxprocs)
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

// TestSubprocessWorker runs inside spawned subprocesses.
// The maxprocs init() has already fired by the time we get here.
func TestSubprocessWorker(t *testing.T) {
	if os.Getenv(subprocessEnvKey) != "1" {
		t.Skip("only runs as subprocess")
	}

	// init() already called Install() — just report the result
	result := subprocessResult{
		GOMAXPROCS: runtime.GOMAXPROCS(0),
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
