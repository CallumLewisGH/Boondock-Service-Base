package models

import "reflect"

func IsEmpty[T any](model T) bool {
	var emptyModel T
	return reflect.DeepEqual(emptyModel, model)
}
