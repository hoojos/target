package endpoint

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	ReqDuration    prometheus.Histogram
	ReqCounter     prometheus.Counter
	OpenConnection prometheus.Gauge
}
