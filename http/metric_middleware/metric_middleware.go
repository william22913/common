package metric_middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/william22913/common/metrics"
)

func NewMetricMiddleware(
	metrics metrics.Metrics,
) MetricMiddleware {
	return &metricMiddleware{
		metrics: metrics,
	}
}

type metricMiddleware struct {
	metrics metrics.Metrics
}

func (m metricMiddleware) Serve(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()

			wr := NewLoggingResponseWriter(w)
			next.ServeHTTP(wr, r)

			if r.URL.Path != "/metrics" {
				m.metrics.GetDefaultMetric().APIHist.WithLabelValues(
					r.URL.Path,
					r.Method,
					strconv.Itoa(wr.statusCode),
				).Observe(float64(time.Since(now).Seconds()))
			}
		},
	)
}
