package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

// Func for build new context with value. Value is a map with resolved services
func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(serviceMap))
	} else {
		return c
	}
}
