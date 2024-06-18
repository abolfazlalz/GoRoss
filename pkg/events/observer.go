package events

type Event[T any] interface {
	Notify(T)
	Subscribe(func(T))
}

type AbstractEventListener[T any] struct {
	listeners []func(T)
}

func NewAbstractEventListener[T any]() *AbstractEventListener[T] {
	return &AbstractEventListener[T]{
		listeners: make([]func(T), 0),
	}
}

func (a *AbstractEventListener[T]) Notify(s T) {
	for _, l := range a.listeners {
		l(s)
	}
}

func (a *AbstractEventListener[T]) Subscribe(f func(T)) {
	a.listeners = append(a.listeners, f)
}
