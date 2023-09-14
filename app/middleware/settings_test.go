package middleware

import (
	"context"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	"github.com/diegosperes/pismo/app/util"
)

func TestAddSettingsContextKey(t *testing.T) {
	var ctxReq context.Context
	request, _ := http.NewRequest(http.MethodGet, "/api/path", nil)

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { ctxReq = r.Context() }
	SettingsMiddleware(util.AppSettings{})(handle)(nil, request, nil)

	assert.NotNil(t, ctxReq)
	assert.IsType(t, ctxReq.Value(SettingsContextKey{}), util.AppSettings{})
}
