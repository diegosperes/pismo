package middleware

import (
	"context"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAddDatabaseContextKey(t *testing.T) {
	var ctxReq context.Context
	request, _ := http.NewRequest(http.MethodGet, "/api/path", nil)

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { ctxReq = r.Context() }
	DatabaseMiddleware(gorm.DB{})(handle)(nil, request, nil)

	assert.NotNil(t, ctxReq)
	assert.IsType(t, ctxReq.Value(DatabaseContextKey{}), gorm.DB{})
}
