package common

import (
	"reflect"
)

func StructToMapClean(v any) map[string]interface{} {
	out := make(map[string]interface{})
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := t.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "-" || jsonTag == "" {
			continue
		}
		key := jsonTag
		if idx := len(key); idx > 0 && key[idx-1] == ',' {
			key = key[:idx-1]
		}
		if field.IsZero() {
			continue
		}
		out[key] = field.Interface()
	}
	return out
}
