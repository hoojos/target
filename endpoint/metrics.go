package endpoint

import "github.com/go-kit/kit/metrics"

type Metrics struct {
	ReqDuration    metrics.Histogram
	ReqCounter     metrics.Counter
	OpenConnection metrics.Gauge
}
