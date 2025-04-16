package models

import "reflect"

func IsEmpty[T any](model any) bool {
	var emptyModel T
	return reflect.DeepEqual(emptyModel, model)
}
