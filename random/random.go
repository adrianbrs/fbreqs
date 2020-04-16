package random

import (
	"math/rand"
	"reflect"
)

// SliceValue returns random value from slice
func SliceValue(slice interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic("SliceValue() given a non-slice type")
	}
	return sliceVal.Index(rand.Intn(sliceVal.Len())).Interface()
}
