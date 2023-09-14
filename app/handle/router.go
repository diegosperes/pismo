package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	middl "github.com/diegosperes/pismo/app/middleware"
	"github.com/diegosperes/pismo/app/util"
)

func GetConfiguredRouter(deps *util.AppDependencies) http.Handler {
	m := middl.NewMiddlewareManager(
		middl.LogRequestMiddleware,
		middl.SettingsMiddleware(deps.Settings),
		middl.DatabaseMiddleware(deps.Database),
	)

	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.NotFound = m.WrapHTTPHandler(HandlerStatusCode(http.StatusNotFound))
	router.MethodNotAllowed = m.WrapHTTPHandler(HandlerStatusCode(http.StatusMethodNotAllowed))

	router.POST("/accounts/", m.Use(CreateAccount))
	router.GET("/accounts/:accountId", m.Use(GetAccount))
	router.POST("/transactions/", m.Use(CreateTransaction))

	return router
}
