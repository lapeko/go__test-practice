package middleware

import (
	"fmt"
	"github.com/lapeko/go__test-practice/2-simple-api/internal/logger"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	xStatusCode int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.xStatusCode = code
	s.ResponseWriter.WriteHeader(code)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		patchedResponse := &statusRecorder{ResponseWriter: res, xStatusCode: http.StatusOK}

		next.ServeHTTP(patchedResponse, req)

		duration := time.Since(start)
		path := req.URL.Path
		if query := req.URL.Query().Encode(); query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}
		logger.Log.WithFields(logrus.Fields{
			"method":   req.Method,
			"path":     path,
			"status":   patchedResponse.xStatusCode,
			"duration": duration.String(),
		}).Info("HTTP Request")
	})
}
