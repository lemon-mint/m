package m

type Result[T any] struct {
	v   T
	err error
}

func Ok[T any](v T) Result[T] {
	return Result[T]{v: v}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Value() (T, error) {
	return r.v, r.err
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("Unwrap called on Err")
	}
	return r.v
}

func (r Result[T]) Match(onOk func(v T), onErr func(err error)) {
	if r.IsOk() {
		onOk(r.v)
	} else {
		onErr(r.err)
	}
}

func RValue[T any](v T, err error) Result[T] {
	return Result[T]{v: v, err: err}
}
