package api_middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
	"go.uber.org/zap"

	"github.com/syunkitada/programming_go/app-sample/pkg/lib/logger"
)

type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	return &captureResponseWriter{w, http.StatusOK}
}

func (lrw *captureResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func CommonHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		tracer := logger.NewTracer()
		tracer.Info("start", zap.String("method", r.Method), zap.String("url", r.URL.String()))
		lrw := NewCaptureResponseWriter(w)

		defer func() {
			elapsed := time.Since(start)
			if err := recover(); err != nil {
				code := http.StatusInternalServerError
				resp, tmpErr := json.Marshal(map[string]string{
					"error": fmt.Sprintf("Internal Server Error: %v", err),
					"code":  "000500", // 予期せぬエラー
				})
				if tmpErr != nil {
					tracer.Error(fmt.Sprintf("failed json.Marshal: err=%v", err))
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(code)
				_, _ = w.Write(resp)
				tracer.Error("end", zap.String("method", r.Method), zap.Int("code", code), zap.String("url", r.URL.String()), zap.Duration("elapsed", elapsed))
				return
			}

			code := lrw.statusCode
			if code >= 500 {
				tracer.Error("end", zap.String("method", r.Method), zap.Int("code", code), zap.String("url", r.URL.String()), zap.Duration("elapsed", elapsed))
			} else if code >= 400 {
				tracer.Warn("end", zap.String("method", r.Method), zap.Int("code", code), zap.String("url", r.URL.String()), zap.Duration("elapsed", elapsed))
			} else {
				tracer.Info("end", zap.String("method", r.Method), zap.Int("code", code), zap.String("url", r.URL.String()), zap.Duration("elapsed", elapsed))
			}
		}()

		next.ServeHTTP(lrw, r)
	})
}
