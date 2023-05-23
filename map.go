package mapper

import "reflect"

func Map2Model[T any](from any) *T {
	from = neverBePtr(from)
	to := reflect.ValueOf(new(T)).Elem()

	if from == nil {
		return nil
	}

	fields := parseField(from)
	for _, f := range fields {
		to = setField(to, f)
	}

	return to.Addr().Interface().(*T)
}

func Map2Models[T any](from []any) []T {
	var to []T
	for _, f := range from {
		to = append(to, *Map2Model[T](f))
	}
	return to
}
