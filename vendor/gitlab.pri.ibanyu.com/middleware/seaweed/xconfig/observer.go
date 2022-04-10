package xconfig

import (
	"context"
	"runtime"
	"sync"

	"github.com/shawnfeng/sutil/slog"
)

// EventHandFunc handle event change
type EventHandFunc func(context.Context, *ChangeEvent)

// ConfigObserver listen config change event
type ConfigObserver struct {
	ch               chan *ChangeEvent
	watchOnce        sync.Once
	applyChangeEvent EventHandFunc
}

// NewConfigObserver ...
func NewConfigObserver(applyChangeEvent EventHandFunc) *ConfigObserver {
	return &ConfigObserver{
		ch:               make(chan *ChangeEvent),
		applyChangeEvent: applyChangeEvent,
	}
}

// HandleChangeEvent send event to listen chan
func (ob *ConfigObserver) HandleChangeEvent(event *ChangeEvent) {
	var changes = map[string]*Change{}
	for k, ce := range event.Changes {
		changes[k] = ce
	}
	if ob.ch == nil {
		slog.Errorf("config observer ch not init")
		return
	}
	event.Changes = changes
	ob.ch <- event
}

// StartWatch watch change event
func (ob *ConfigObserver) StartWatch(ctx context.Context) {
	fun := "ConfigObserver Watch"
	if ob.ch == nil {
		slog.Errorf("%s config observer ch not init", fun)
		return
	}
	ob.watchOnce.Do(func() {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 4096)
					buf = buf[:runtime.Stack(buf, false)]
					slog.Errorf("%s recover err: %v, stack: %s", fun, err, string(buf))
				}
			}()
			for {
				select {
				case <-ctx.Done():
					slog.Infof("%s context done err:%v", fun, ctx.Err())
					return
				case ce, ok := <-ob.ch:
					if !ok {
						slog.Infof("%s change event channel closed", fun)
						return
					}
					ob.applyChangeEvent(ctx, ce)
				}
			}
		}()
	})
}
