package middleware

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestWrapHandle(t *testing.T) {
	handleCalled := false
	middlewareCalled := false

	middleware := func(handle httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			middlewareCalled = true
			handle(w, r, ps)
		}
	}

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { handleCalled = true }

	manager := NewMiddlewareManager(middleware)
	request, _ := http.NewRequest(http.MethodGet, "/api/path", nil)
	manager.Use(handle)(nil, request, nil)

	assert.Equal(t, true, middlewareCalled)
	assert.Equal(t, true, handleCalled)
}

func TestWrapHTTPHandler(t *testing.T) {
	handleCalled := false
	middlewareCalled := false

	middleware := func(handle httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			middlewareCalled = true
			handle(w, r, ps)
		}
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleCalled = true })

	manager := NewMiddlewareManager(middleware)
	request, _ := http.NewRequest(http.MethodGet, "/api/path", nil)
	manager.WrapHTTPHandler(handler).ServeHTTP(nil, request)

	assert.Equal(t, true, middlewareCalled)
	assert.Equal(t, true, handleCalled)
}
