package agent

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/utils/guid"
)

func TestValidateDeployment(t *testing.T) {
	cases := []struct {
		name       string
		deployment string
		valid      bool
	}{
		{"empty", "", true},
		{"alphanumeric", "production", true},
		{"hyphen and dot", "prod-us.v2", true},
		{"colon", "prod:us", true},
		{"slash", "a/b", true},
		{"non-ascii", "prodüction", true},
		{"max length", strings.Repeat("a", MaxDeploymentLength), true},

		{"underscore reserved", "prod_us", false},
		{"space", "prod us", false},
		{"tab", "prod\tus", false},
		{"newline", "prod\nus", false},
		{"del", "prod\x7fus", false},
		{"nul", "prod\x00us", false},
		{"too long", strings.Repeat("a", MaxDeploymentLength+1), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateDeployment(c.deployment)
			if c.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestValidateThreadID(t *testing.T) {
	require.NoError(t, ValidateThreadID(""))
	require.NoError(t, ValidateThreadID(guid.New(guid.AgentThreadPrefix)))
	for _, id := range []string{"AT_", "AT_not-valid", "AT_0123456789ab", "RM_" + guid.New("")} {
		require.Error(t, ValidateThreadID(id), id)
	}
}
