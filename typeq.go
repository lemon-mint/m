package m

import "reflect"

type TypeKind uint8

const (
	TypeKindUnknown TypeKind = iota
	TypeKindStructLike
	TypeKindArrayLike
)

type TypeInfo struct {
	Kind TypeKind
}

func TypeQ[T any]() (TI TypeInfo) {
	zv := Zero[T]()
	rv := reflect.ValueOf(zv)
	rt := rv.Type()
	for rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}

	if rt.Kind() == reflect.Struct {
		TI.Kind = TypeKindStructLike
	} else if rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice {
		TI.Kind = TypeKindArrayLike
	} else {
		TI.Kind = TypeKindUnknown
	}

	return
}
