package xutil

import (
	"context"
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

// Future safe goroutine
type Future struct {
	ch  chan struct{}
	fn  Func
	err error
}

// Func func of future
type Func func() error

// NewFuture create a future
func NewFuture(fn Func) *Future {
	f := &Future{
		ch: make(chan struct{}),
		fn: fn,
	}
	f.start()
	return f
}

func (f *Future) start() {
	go func() {
		defer func() {
			if ret := recover(); ret != nil {
				if err, ok := ret.(error); ok {
					f.err = pkgErrors.WithStack(err)
				} else {
					retStr := fmt.Sprint(ret)
					f.err = pkgErrors.WithStack(errors.New(retStr))
				}
			}

			close(f.ch)
		}()

		// 执行结果
		f.err = f.fn()
		return
	}()
}

// Get return result of future
func (f *Future) Get(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-f.ch:
		return f.err
	}
}

// Done check if future done
func (f *Future) Done() bool {
	select {
	case <-f.ch:
		return true
	default:
		return false
	}
}

// WaitAllFutures wait all futures done
func WaitAllFutures(ctx context.Context, futures ...*Future) {
	for _, feature := range futures {
		feature.Get(ctx)
	}
}
