package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

// Use Middleware go pattern to manage request lifecycle.
type MiddlewareManager struct {
	Middlewares []Middleware
}

func NewMiddlewareManager(m ...Middleware) *MiddlewareManager {
	return &MiddlewareManager{Middlewares: m}
}

func (mm *MiddlewareManager) Use(handler httprouter.Handle) httprouter.Handle {
	for _, middleware := range mm.Middlewares {
		handler = middleware(handler)
	}

	return handler
}

func (mm *MiddlewareManager) WrapHTTPHandler(handler http.Handler) http.Handler {
	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { handler.ServeHTTP(w, r) }

	for _, middleware := range mm.Middlewares {
		handle = middleware(handle)
	}

	h := func(w http.ResponseWriter, r *http.Request) { handle(w, r, httprouter.ParamsFromContext(r.Context())) }
	return http.HandlerFunc(h)
}
