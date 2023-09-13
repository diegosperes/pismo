package middleware

import (
	"log"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type logTestWritter struct {
	Output [][]byte
}

func (w *logTestWritter) Write(p []byte) (n int, err error) {
	w.Output = append(w.Output, p)
	return len(p), nil
}

func TestLogRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/api/path", nil)

	logWritter := &logTestWritter{}
	log.SetOutput(logWritter)

	called := false
	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { called = true }
	LogRequestMiddleware(handle)(nil, request, nil)

	assert.Equal(t, true, called)
	assert.Contains(t, string(logWritter.Output[0]), "INFO Request method=GET path=/api/path")
}
