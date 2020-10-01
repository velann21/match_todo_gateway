package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Metrics() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			handler.ServeHTTP(w, req)
			return
		})
	}
}