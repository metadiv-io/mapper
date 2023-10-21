package types

import "reflect"

type Field struct {
	Name  string
	Type  reflect.Type
	Value reflect.Value
}
