package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logrus.Infof("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
