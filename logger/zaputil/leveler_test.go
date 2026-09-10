package zaputil

import (
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

type componentLevels map[string]zapcore.Level

func (c componentLevels) ResolveComponentLevel(component string) (zapcore.Level, bool) {
	lvl, ok := c[component]
	return lvl, ok
}

func discardRoot(lvl zapcore.Level) *ComponentLeveler {
	return NewRootComponentLeveler(zapcore.AddSync(io.Discard), FixedComponentLevel(lvl))
}

func TestComponentLevelerMemoizes(t *testing.T) {
	root := discardRoot(zapcore.InfoLevel)

	require.Same(t, root.WriteEnabler("sub"), root.WriteEnabler("sub"))
	require.NotSame(t, root.WriteEnabler("sub"), root.WriteEnabler("other"))
	require.Equal(t, root.ComponentLevel("sub"), root.ComponentLevel("sub"))
}

func TestComponentLevelerWidensParent(t *testing.T) {
	root := discardRoot(zapcore.InfoLevel)
	child := NewComponentLeveler(root, componentLevels{"sub": zapcore.DebugLevel})

	require.True(t, child.ComponentLevel("sub").Enabled(zapcore.DebugLevel))
	require.False(t, child.ComponentLevel("other").Enabled(zapcore.DebugLevel))
	require.True(t, child.ComponentLevel("other").Enabled(zapcore.InfoLevel))
}

func TestComponentLevelerCannotQuietParent(t *testing.T) {
	root := discardRoot(zapcore.DebugLevel)
	child := NewComponentLeveler(root, componentLevels{"sub": zapcore.ErrorLevel})

	require.True(t, child.ComponentLevel("sub").Enabled(zapcore.DebugLevel))
}

func TestComponentLevelerRefresh(t *testing.T) {
	root := discardRoot(zapcore.InfoLevel)
	levels := componentLevels{}
	child := NewComponentLeveler(root, levels)

	enab := child.ComponentLevel("sub")
	require.False(t, enab.Enabled(zapcore.DebugLevel))

	levels["sub"] = zapcore.DebugLevel
	child.Refresh()
	require.True(t, enab.Enabled(zapcore.DebugLevel))

	delete(levels, "sub")
	child.Refresh()
	require.False(t, enab.Enabled(zapcore.DebugLevel))
	require.True(t, enab.Enabled(zapcore.InfoLevel))
}

// Refresh walks the entry map while other goroutines materialize into it.
func TestComponentLevelerConcurrentRefresh(t *testing.T) {
	root := discardRoot(zapcore.InfoLevel)
	child := NewComponentLeveler(root, FixedComponentLevel(zapcore.DebugLevel))

	var wg sync.WaitGroup
	for i := range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range 200 {
				c := fmt.Sprintf("c%d.s%d", i, j%16)
				if child.WriteEnabler(c) == nil || root.ComponentLevel(c) == nil {
					t.Error("nil enabler")
					return
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 200 {
			root.Refresh()
			child.Refresh()
		}
	}()

	wg.Wait()
}
