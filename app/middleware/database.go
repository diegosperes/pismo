package middleware

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type DatabaseContextKey struct{}

func DatabaseMiddleware(db gorm.DB) Middleware {
	return func(handle httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := r.Context()
			ctxDB := context.WithValue(ctx, DatabaseContextKey{}, db)
			handle(w, r.WithContext(ctxDB), ps)
		}
	}
}
