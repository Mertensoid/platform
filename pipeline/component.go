package pipeline

import (
	"net/http"
)

// Data for request processing
type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

// Set error
func (mwc *ComponentContext) Error(err error) {
	mwc.error = err
}

// Get error
func (mwc *ComponentContext) GetError() error {
	return mwc.error
}

// Middleware interface
type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentContext, next func(*ComponentContext))
}
