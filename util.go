package mapper

import (
	"fmt"
	"reflect"
)

type field struct {
	Name  string
	Type  reflect.Type
	Value reflect.Value
}

func neverBePtr(v any) any {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return reflect.ValueOf(v).Elem().Interface()
	}
	return v
}

func parseField(v any) []field {
	v = neverBePtr(v)
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		panic("only struct is supported")
	}
	var fields []field
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			fields = append(fields, parseField(reflect.ValueOf(v).FieldByName(f.Name).Interface())...)
			continue
		}
		if f.Type.Kind() == reflect.Ptr {
			f.Type = f.Type.Elem()
		}
		if !f.IsExported() {
			continue
		}
		switch f.Type.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan, reflect.Interface,
			reflect.String,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Bool:
			fields = append(fields, field{
				Name:  f.Name,
				Type:  f.Type,
				Value: reflect.ValueOf(v).Field(i),
			})
		default:
			continue
		}
	}
	return fields
}

func setField(v reflect.Value, field field) reflect.Value {
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if f.Anonymous {
			v.Field(i).Set(setField(v.Field(i), field))
			continue
		}
		if f.Name == field.Name {
			_, ok := v.Type().FieldByName(field.Name)
			if ok {
				fmt.Println(field.Type, f.Type)
				if field.Type == f.Type {
					v.FieldByName(field.Name).Set(field.Value)
				}
			}
		}
	}
	return v
}
