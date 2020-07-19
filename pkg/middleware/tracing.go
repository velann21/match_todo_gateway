package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/google/uuid"
	"github.com/velann21/match_todo_gateway_srv/pkg/permissions"
)

func TraceLogger() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			eventType := permissions.EventsPermission[req.URL.Path]
			if eventType == "" {
				handler.ServeHTTP(w, req)
				return
			}
			syntheticTraceID := uuid.New().String()
			req.Header.Add("X-TraceID", syntheticTraceID)
			req.Header.Add("X-EventType", eventType)
			handler.ServeHTTP(w, req)
			return
		})
	}
}