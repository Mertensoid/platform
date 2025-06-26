package services

import (
	"context"
	"reflect"
	"sync"
)

// Adds service with transient lifecycle
func AddTransient(factoryFunc interface{}) (err error) {
	return addService(Transient, factoryFunc)
}

// Adds service with scoped lifecycle
func AddScoped(factoryFunc interface{}) (err error) {
	return addService(Scoped, factoryFunc)
}

// Adds service with scoped lifecycle. Checks that service will be called just once.
func AddSingleton(factoryFunc interface{}) (err error) {
	factoryFuncVal := reflect.ValueOf(factoryFunc)
	if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
		var results []reflect.Value
		once := sync.Once{}
		wrapper := reflect.MakeFunc(factoryFuncVal.Type(),
			func([]reflect.Value) []reflect.Value {
				once.Do(func() {
					results = invokeFunction(context.TODO(), factoryFuncVal)
				})
				return results
			})
		err = addService(Singleton, wrapper.Interface())
	}
	return
}
