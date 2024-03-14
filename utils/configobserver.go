package utils

import (
	"container/list"
	"fmt"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"

	"github.com/livekit/protocol/logger"
)

type ConfigBuilder[T any] interface {
	New() (*T, error)
}

type ConfigDefaulter[T any] interface {
	InitDefaults(*T) error
}

type ConfigObserver[T any] struct {
	builder ConfigBuilder[T]
	watcher *fsnotify.Watcher
	mu      sync.Mutex
	cbs     list.List
}

func NewConfigObserver[T any](path string, builder ConfigBuilder[T]) (*ConfigObserver[T], *T, error) {
	c := &ConfigObserver[T]{
		builder: builder,
	}

	config, err := c.load(path)
	if err != nil {
		return nil, nil, err
	}

	if path != "" {
		c.watcher, err = fsnotify.NewWatcher()
		if err != nil {
			return nil, nil, err
		}
		if err := c.watcher.Add(path); err != nil {
			c.watcher.Close()
			return nil, nil, err
		}
		go c.watch()
	}

	return c, config, nil
}

func (c *ConfigObserver[T]) Close() {
	if c != nil && c.watcher != nil {
		c.watcher.Close()
	}
}

func (c *ConfigObserver[T]) EmitConfigUpdate(conf *T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for e := c.cbs.Front(); e != nil; e = e.Next() {
		e.Value.(func(*T))(conf)
	}
}

func (c *ConfigObserver[T]) Observe(cb func(*T)) func() {
	if c == nil {
		return func() {}
	}
	c.mu.Lock()
	e := c.cbs.PushBack(cb)
	c.mu.Unlock()

	return func() {
		c.mu.Lock()
		c.cbs.Remove(e)
		c.mu.Unlock()
	}
}

func (c *ConfigObserver[T]) watch() {
	for {
		select {
		case event, ok := <-c.watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Remove) {
				if err := c.watcher.Add(event.Name); err != nil {
					logger.Errorw("unable to rewatch config file", err, "file", event.Name)
				}
			}
			if event.Has(fsnotify.Write | fsnotify.Remove) {
				if err := c.reload(event.Name); err != nil {
					logger.Errorw("unable to update config file", err, "file", event.Name)
				} else {
					logger.Infow("config file has been updated", "file", event.Name)
				}
			}
		case err, ok := <-c.watcher.Errors:
			if !ok {
				return
			}
			logger.Errorw("config file watcher error", err)
		}
	}
}

func (c *ConfigObserver[T]) reload(path string) error {
	conf, err := c.load(path)
	if err != nil {
		return err
	}

	c.EmitConfigUpdate(conf)
	return nil
}

func (c *ConfigObserver[T]) load(path string) (*T, error) {
	conf, err := c.builder.New()
	if err != nil {
		return nil, err
	}

	if path != "" {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		if len(b) == 0 {
			return nil, fmt.Errorf("cannot parse config: file empty")
		}

		if err := yaml.Unmarshal(b, conf); err != nil {
			return nil, fmt.Errorf("cannot parse config: %v", err)
		}
	}

	if d, ok := c.builder.(ConfigDefaulter[T]); ok {
		d.InitDefaults(conf)
	}

	return conf, err
}
