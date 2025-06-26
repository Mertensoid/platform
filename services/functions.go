package services

import (
	"context"
	"errors"
	"reflect"
)

// Uses CallForContext with background context
func Call(target interface{}, otherArgs ...interface{}) ([]interface{}, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

// Calls function using services for getting values
func CallForContext(c context.Context, target interface{},
	otherArgs ...interface{}) (result []interface{}, err error) {
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() == reflect.Func {
		resultVals := invokeFunction(c, targetVal, otherArgs...)
		result = make([]interface{}, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			result[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("Only functions can be invoked")
	}
	return
}
