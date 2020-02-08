package logging

import (
	"bookshelf_service/src/logger"
	"net/http"
)

func LoggingMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Infow("incoming request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
