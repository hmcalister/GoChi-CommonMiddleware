package commonMiddleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// Perform logging of requests using a slog logger.
//
// Should be placed early if not first in middleware stack.
func SlogLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		wrappedWriter := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		requestTimeReceived := time.Now()

		next.ServeHTTP(wrappedWriter, r)

		requestTimeResolved := time.Now()
		slog.Info("SlogLoggerMiddleware",
			"URL", r.URL.Path,
			"Protocol", r.Proto,
			"RemoteIP", r.RemoteAddr,
			"Status", wrappedWriter.Status(),
			"UserAgent", r.Header.Get("User-Agent"),
			"Latency_ms", float32(requestTimeResolved.Sub(requestTimeReceived).Nanoseconds()/1_000_000.0),
			"BytesReceived", r.ContentLength,
			"BytesSent", wrappedWriter.BytesWritten(),
		)
	})
}
