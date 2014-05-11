package helpers

import (
	"reflect"
)

func StructToBSONMap(st interface{}) (m map[string]interface{}) {

	s := reflect.ValueOf(&st).Elem()
	typeOfT := s.Type()

	m = make(map[string]interface{})

	for i := 0; i < s.NumField(); i++ {

		field := s.Field(i)
		typeField := typeOfT.Field(i)

		fieldName := typeField.Tag.Get("map")
		if fieldName == "" {

			fieldName = typeField.Name
		}

		m[fieldName] = field.Interface()
	}

	return
}

func IsNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}
