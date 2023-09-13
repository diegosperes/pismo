package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	m "github.com/diegosperes/pismo/app/middleware"
)

func GetConfiguredRouter() http.Handler {
	middl := m.NewMiddlewareManager(m.LogRequestMiddleware)

	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.NotFound = middl.WrapHTTPHandler(HandlerStatusCode(http.StatusNotFound))
	router.MethodNotAllowed = middl.WrapHTTPHandler(HandlerStatusCode(http.StatusMethodNotAllowed))

	router.POST("/accounts/", middl.Use(CreateAccount))
	router.GET("/accounts/:accountId", middl.Use(GetAccount))
	router.POST("/transactions/", middl.Use(CreateTransaction))

	return router
}
