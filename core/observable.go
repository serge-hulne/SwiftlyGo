package core

import (
	"gocore/reactive"
	"sync"
)

type internalObservable interface {
	addSubscriber(func())
}

type Observable[T any] struct {
	value     T
	listeners []func(T)
	mu        sync.Mutex
}

type ReadonlyObservable[T any] interface {
	Get() T
	Subscribe(func(T))
}

func NewObservable[T any](initial T) *Observable[T] {
	return &Observable[T]{value: initial}
}

func (o *Observable[T]) Get() T {
	trackObservable(o)
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.value
}

func (o *Observable[T]) Set(v T) {
	o.mu.Lock()
	o.value = v
	for _, listener := range o.listeners {
		listener(v)
	}
	o.mu.Unlock()
}

func (o *Observable[T]) Subscribe(listener func(T)) {
	o.mu.Lock()
	o.listeners = append(o.listeners, listener)
	listener(o.value)
	o.mu.Unlock()
}

func (o *Observable[T]) addSubscriber(update func()) {
	o.mu.Lock()
	o.listeners = append(o.listeners, func(T) { update() })
	o.mu.Unlock()
}

var _ reactive.ReadonlyObservable[any] = (*Observable[any])(nil)
