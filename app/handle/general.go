package handle

import (
	"net/http"

	"github.com/diegosperes/pismo/app/util"
)

func HandlerStatusCode(statusCode int) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			util.WriteJsonError(w, statusCode)
		},
	)
}
