package reactive

type ReadonlyObservable[T any] interface {
	Get() T
	Subscribe(func(T))
}
