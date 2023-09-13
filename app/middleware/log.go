package middleware

import (
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LogRequestMiddleware(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		slog.Info("Request", "method", r.Method, "path", r.URL.Path)
		handle(w, r, ps)
	}
}
