package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) writeHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String())

		fmt.Printf(
			"[Performance]: URL=%s; Method=%s; Status=%d; Duration=%v\n",
			r.URL,
			r.Method,
			wrappedWriter.status,
			duration,
		)
	})
}
