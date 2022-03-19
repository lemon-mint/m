package m

type Option[T any] struct {
	v    T
	none bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{v: v, none: false}
}

func None[T any]() Option[T] {
	return Option[T]{none: true}
}

func (o Option[T]) IsNone() bool {
	return o.none
}

func (o Option[T]) IsSome() bool {
	return !o.none
}

func (o Option[T]) Value() (T, bool) {
	return o.v, o.IsSome()
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("Unwrap called on None")
	}
	return o.v
}

func (o Option[T]) Match(onSome func(v T), onNone func()) {
	if o.IsSome() {
		onSome(o.v)
	} else {
		onNone()
	}
}
