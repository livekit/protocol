// Package leaftest asserts that a package's dependency closure stops where its
// author said it does.
//
// A package that promises to import nothing but the standard library is making a
// promise its consumers rely on: reading a value out of it must not drag in the
// machinery that produced the value. Nothing in the language enforces that, and
// the import that breaks it is one keystroke and no review comment away. Assert
// turns the promise into a test.
package leaftest

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// Assert fails t unless the package in the current directory imports nothing
// beyond the standard library and the import paths in allowed.
//
// Call it from a test in the package being guarded; the working directory is what
// names the package, so no import path is passed in and none can go stale.
//
// Test-only imports are not counted: `go list -deps` reports what the package
// itself needs, so a guarded package's own test may use testify, this helper, or
// anything else without weakening what the guard means.
//
// Naming a path in allowed is the decision this is here to make visible. Keep the
// list in one place the reviewer of a widening diff will see.
func Assert(t testing.TB, allowed ...string) {
	t.Helper()
	if err := check(".", allowed); err != nil {
		t.Fatal(err)
	}
}

// check is Assert's whole decision, separated so it can be tested: testing.TB
// cannot be implemented outside the testing package, so a helper that only ever
// fails a *testing.T cannot be shown to fail.
func check(dir string, allowed []string) error {
	self, err := list(dir, "-f", "{{.ImportPath}}")
	if err != nil {
		return err
	}
	if len(self) != 1 {
		return fmt.Errorf("leaftest: expected one package in %s, got %d: %v", dir, len(self), self)
	}

	permitted := map[string]bool{self[0]: true}
	for _, p := range allowed {
		permitted[p] = true
	}

	// The package names itself, so there is no constant here to fall out of date
	// with a rename and leave the check passing against a package that no longer
	// exists. go list -deps always reports the package among its own deps, so an
	// empty result is a failure list already returned above, not a silent pass.
	deps, err := list(dir, "-deps")
	if err != nil {
		return err
	}

	var external []string
	for _, p := range deps {
		if !permitted[p] && !isStdlib(p) {
			external = append(external, p)
		}
	}
	if len(external) == 0 {
		return nil
	}
	return fmt.Errorf("leaftest: %s must import only the standard library%s, but its dependency "+
		"closure reaches %d package(s):\n\t%s\n\nSee the rule in that package's doc comment. "+
		"If the new dependency is intended, add it to the allow list this test passes.",
		self[0], andAllowed(allowed), len(external), strings.Join(external, "\n\t"))
}

func andAllowed(allowed []string) string {
	if len(allowed) == 0 {
		return ""
	}
	return " and " + strings.Join(allowed, ", ")
}

// isStdlib reports whether an import path belongs to the standard library. A
// stdlib path has no dot in its first segment, because a module path starts with
// a domain.
func isStdlib(pkg string) bool {
	first, _, _ := strings.Cut(pkg, "/")
	return !strings.Contains(first, ".")
}

func list(dir string, args ...string) ([]string, error) {
	cmd := exec.Command("go", append(append([]string{"list"}, args...), ".")...)
	cmd.Dir = dir
	out, err := cmd.Output()

	var exit *exec.ExitError
	if errors.As(err, &exit) {
		return nil, fmt.Errorf("leaftest: go list %s . in %s: %w\n%s", strings.Join(args, " "), dir, err, exit.Stderr)
	}
	if err != nil {
		return nil, fmt.Errorf("leaftest: go list %s . in %s: %w", strings.Join(args, " "), dir, err)
	}
	return strings.Fields(string(out)), nil
}
