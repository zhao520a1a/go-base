package util

import (
	"fmt"
	"reflect"
)

// isZero reports whether a value is a zero value
// Including support: Bool, Array, String, Float32, Float64, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr
// Map, Slice, Interface, Struct
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Map, reflect.Slice:
		return v.IsNil() || v.Len() == 0
	case reflect.Interface:
		return v.IsNil()
	case reflect.Invalid:
		return true
	}

	if v.Kind() != reflect.Struct {
		return false
	}

	// Traverse the Struct and only return true
	// if all of its fields return IsZero == true
	n := v.NumField()
	for i := 0; i < n; i++ {
		vf := v.Field(i)
		if !isZero(vf) {
			return false
		}
	}
	return true
}

// OmitEmpty is a helper function to clear where map zero value
func OmitEmpty(where map[string]interface{}, omitKey []string) map[string]interface{} {
	for _, key := range omitKey {
		v, ok := where[key]
		if !ok {
			continue
		}

		if isZero(reflect.ValueOf(v)) {
			delete(where, key)
		}
	}
	return where
}

func main() {
	where := map[string]interface{}{
		"score": 0,
		"age":   35,
	}
	finalWhere := OmitEmpty(where, []string{"score", "age"})
	fmt.Printf("%v", finalWhere)
}
