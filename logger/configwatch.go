package logger

import (
	"bytes"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	// ConfigPathEnv names a file holding the same keys as the service config's `logging` block.
	// Point it inside a mounted ConfigMap to change levels without restarting: kubelet refreshes
	// the mount in place, and the next poll pushes the new values into the live logger.
	ConfigPathEnv = "LK_LOG_CONFIG_PATH"
	// ConfigIntervalEnv overrides the poll interval as a Go duration (e.g. "10s").
	ConfigIntervalEnv = "LK_LOG_CONFIG_INTERVAL"

	defaultConfigWatchInterval = 30 * time.Second
)

var configWatchOnce sync.Once

// startConfigWatchFromEnv wires the watcher for the first logger the process builds, which is the
// one whose Config the service keeps. Every binary that uses this package reaches it through
// newSharedConfig, so none of them need their own flag or call site.
func startConfigWatchFromEnv(conf *Config) {
	path := os.Getenv(ConfigPathEnv)
	if path == "" {
		return
	}
	interval := defaultConfigWatchInterval
	if v := os.Getenv(ConfigIntervalEnv); v != "" {
		if d, err := time.ParseDuration(v); err == nil && d > 0 {
			interval = d
		}
	}
	configWatchOnce.Do(func() {
		WatchConfigFile(conf, path, interval)
	})
}

// WatchConfigFile applies path to conf every interval until the returned stop is called.
//
// Polling rather than fsnotify on purpose: a ConfigMap volume update swaps the `..data` symlink
// instead of rewriting the file, so a watch on the file itself never fires.
func WatchConfigFile(conf *Config, path string, interval time.Duration) (stop func()) {
	done := make(chan struct{})
	// The config the process started with. Every file is applied over this, never over whatever
	// the previous file left in force, so an empty file restores the startup levels and a
	// component_levels entry that disappears from the file stops applying.
	baseline := conf.snapshot()
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		var last []byte
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if applied, ok := applyConfigFile(conf, baseline, path, last); ok {
					last = applied
				}
			}
		}
	}()

	var once sync.Once
	return func() { once.Do(func() { close(done) }) }
}

// applyConfigFile decodes path over baseline and pushes the result into conf when the bytes differ
// from last, returning the bytes it applied. ok is false when nothing was applied (unreadable,
// unchanged or invalid), and the config already in force stays untouched.
//
// Decoding over baseline rather than over conf is what makes the file a declarative overlay: keys
// it omits fall back to the startup values instead of inheriting the previous file's.
func applyConfigFile(conf, baseline *Config, path string, last []byte) (applied []byte, ok bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		// An optional ConfigMap that is not mounted yet is the normal steady state, not something
		// to log once per interval forever. A file that goes away is deliberately not treated as a
		// reset either: a transient read error would otherwise flap levels. Emptying the file to
		// `{}` is the reset.
		return nil, false
	}
	if bytes.Equal(data, last) {
		return nil, false
	}

	next := baseline.snapshot()
	if err := yaml.Unmarshal(data, next); err != nil {
		Warnw("could not parse log config, keeping the one in force", err, "path", path)
		return nil, false
	}
	if err := conf.Update(next); err != nil {
		Warnw("could not apply log config, keeping the one in force", err, "path", path)
		return nil, false
	}
	return data, true
}
