package core

import (
	"context"
	"sync"
)

// === Derived (reactive/computed) ===

type Derived[T any] struct {
	mu     sync.Mutex
	value  T
	subs   []func(T)
	cancel context.CancelFunc
}

// Derive creates a reactive computed observable.
func Derive[T any](compute func() T) *Derived[T] {
	d := &Derived[T]{}

	ctx, cancel := context.WithCancel(context.Background())
	d.cancel = cancel

	var deps []internalObservable
	depMu := sync.Mutex{}

	tracker := &dependencyTracker{
		add: func(obs internalObservable) {
			depMu.Lock()
			deps = append(deps, obs)
			depMu.Unlock()
		},
	}

	result := runWithTracker(ctx, tracker, compute)
	d.value = result

	update := func() {
		newVal := compute()
		d.mu.Lock()
		d.value = newVal
		for _, sub := range d.subs {
			go sub(newVal) // async dispatch
		}
		d.mu.Unlock()
	}

	for _, dep := range deps {
		dep.addSubscriber(func() {
			go update()
		})
	}

	return d
}

func (d *Derived[T]) Get() T {
	d.mu.Lock()
	val := d.value
	d.mu.Unlock()
	return val
}

func (d *Derived[T]) Subscribe(sub func(T)) {
	d.mu.Lock()
	d.subs = append(d.subs, sub)
	val := d.value
	d.mu.Unlock()
	sub(val)
}

func Map[T any, U any](d *Derived[T], f func(T) U) *Derived[U] {
	return Derive(func() U {
		return f(d.Get())
	})
}

// === Dependency Tracking Context ===

type contextKey string

var trackerKey = contextKey("reactive")

type dependencyTracker struct {
	add func(internalObservable)
}

func runWithTracker[T any](ctx context.Context, tracker *dependencyTracker, compute func() T) T {
	prevCtxMu.Lock()
	prev := currentCtx
	currentCtx = context.WithValue(ctx, trackerKey, tracker)
	prevCtxMu.Unlock()

	val := compute()

	prevCtxMu.Lock()
	currentCtx = prev
	prevCtxMu.Unlock()

	return val
}

var (
	currentCtx context.Context
	prevCtxMu  sync.Mutex
)

// ==== Observable hook ====

func trackObservable(obs internalObservable) {
	prevCtxMu.Lock()
	ctx := currentCtx
	prevCtxMu.Unlock()

	if ctx != nil {
		if tracker, ok := ctx.Value(trackerKey).(*dependencyTracker); ok && tracker != nil {
			tracker.add(obs)
		}
	}
}
