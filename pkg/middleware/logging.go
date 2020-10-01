package middleware

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logging() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			logrus.Info("Inside Logging")
			start := time.Now()
			route := req.URL.Path
			logrus.Info("start time ----",route,"-------->>>>>>>>", start)
			handler.ServeHTTP(w, req)
			logrus.Info("End time  ----",route,"-------->>>>>>>>", time.Now())
			logrus.Info("Time took to execute  time ----",route,"-------->>>>>>>>", time.Since(start))
			return
		})
	}
}
