package middleware

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/diegosperes/pismo/app/util"
)

type SettingsContextKey struct{}

func SettingsMiddleware(settings util.AppSettings) Middleware {
	return func(handle httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := r.Context()
			ctxSettings := context.WithValue(ctx, SettingsContextKey{}, settings)
			handle(w, r.WithContext(ctxSettings), ps)
		}
	}
}
