package mapper

import (
	"reflect"

	"github.com/metadiv-io/mapper/internal/util"
)

// Map2Model converts model to model.
// The model must be a struct.
// The target model set as generic type.
func Map2Model[T any](from any) *T {
	from = util.NeverBePtr(from)
	to := reflect.ValueOf(new(T)).Elem()

	if from == nil {
		return nil
	}

	fields := util.ParseField(from)
	for _, f := range fields {
		to = util.SetField(to, f)
	}

	return to.Addr().Interface().(*T)
}

// Map2Models converts models to models.
// The model must be a struct.
// The target model set as generic type.
func Map2Models[T any](from []any) []T {
	var to []T
	for _, f := range from {
		to = append(to, *Map2Model[T](f))
	}
	return to
}

// MapModel2Model converts model to model.
// The model must be a struct.
// The target model set as generic type.
func MapModel2Model[T any](from any, to *T) *T {
	from = util.NeverBePtr(from)

	if from == nil || to == nil {
		return nil
	}

	new := reflect.ValueOf(new(T)).Elem()

	fields := util.ParseField(to)
	for _, f := range fields {
		new = util.SetField(new, f)
	}

	fields = util.ParseField(from)
	for _, f := range fields {
		new = util.SetField(new, f)
	}

	return new.Addr().Interface().(*T)
}

// MapModels2Models converts models to models.
// The model must be a struct.
// The target model set as generic type.
func MapModels2Models[T any](from []any, to []T) []T {
	for i, f := range from {
		to[i] = *MapModel2Model[T](f, &to[i])
	}
	return to
}
