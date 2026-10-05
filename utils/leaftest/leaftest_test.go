package leaftest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A guarded package's test imports this one, so it must be a leaf itself.
func TestPackageIsALeaf(t *testing.T) {
	Assert(t)
}

func TestAcceptsAStdlibOnlyPackage(t *testing.T) {
	require.NoError(t, check("../must", nil))
}

func TestRejectsAPackageThatReachesFurther(t *testing.T) {
	err := check("../configutil", nil)
	require.Error(t, err)

	msg := err.Error()
	assert.Contains(t, msg, "github.com/livekit/protocol/utils/configutil",
		"the failure must name the package")
	assert.Contains(t, msg, "fsnotify", "and what it reached")
	assert.Contains(t, msg, "doc comment", "and where the rule is written")
}

// Allowing a path subtracts it from the report; allowing everything a package
// reaches makes it pass.
func TestAllowListSubtracts(t *testing.T) {
	bare := check("../configutil", nil)
	require.Error(t, bare)

	reached := reachedPackages(t, bare)
	require.Greater(t, len(reached), 1, "need more than one dep for this to mean anything")

	narrowed := check("../configutil", reached[:1])
	require.Error(t, narrowed)
	assert.Len(t, reachedPackages(t, narrowed), len(reached)-1)

	require.NoError(t, check("../configutil", reached),
		"allowing everything it reaches is what makes a non-leaf pass")
}

// An allow list may name a path the package does not import.
func TestAllowingSomethingUnusedIsNotAnError(t *testing.T) {
	require.NoError(t, check("../must", []string{"github.com/nothing/here"}))
}

func TestReportsAMissingDirectory(t *testing.T) {
	err := check("../does-not-exist", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "go list")
}

func TestIsStdlib(t *testing.T) {
	for _, p := range []string{"errors", "os/exec", "crypto/x509", "slices"} {
		assert.True(t, isStdlib(p), p)
	}
	for _, p := range []string{
		"github.com/livekit/protocol/utils",
		"golang.org/x/sync/errgroup",
		"gopkg.in/yaml.v3",
	} {
		assert.False(t, isStdlib(p), p)
	}
}

// reachedPackages pulls the offending import paths out of a failure, so a test
// can assert on the set and not the prose around it.
func reachedPackages(t *testing.T, err error) []string {
	t.Helper()
	var out []string
	for _, line := range strings.Split(err.Error(), "\n") {
		if p, ok := strings.CutPrefix(line, "\t"); ok {
			out = append(out, p)
		}
	}
	require.NotEmpty(t, out, "could not parse the failure: %v", err)
	return out
}
