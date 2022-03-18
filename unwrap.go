package m

func Unwrap[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func Unwrap2[T0 any, T1 any](v0 T0, v1 T1, err error) (T0, T1) {
	if err != nil {
		panic(err)
	}
	return v0, v1
}

func Unwrap3[T0 any, T1 any, T2 any](v0 T0, v1 T1, v2 T2, err error) (T0, T1, T2) {
	if err != nil {
		panic(err)
	}
	return v0, v1, v2
}

func Unwrap4[T0 any, T1 any, T2 any, T3 any](v0 T0, v1 T1, v2 T2, v3 T3, err error) (T0, T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return v0, v1, v2, v3
}
