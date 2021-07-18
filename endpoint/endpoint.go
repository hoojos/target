package endpoint

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"time"
)

type Endpoint struct {
	method  string
	URL     string
	headers map[string]string
	body    []byte
	latency time.Duration
	metric  Metrics
}

func (e Endpoint) Handle(ctx *routing.Context) error {
	time.Sleep(e.latency)
	for k, v := range e.headers {
		ctx.Response.Header.Set(k, v)
	}
	_, err := ctx.Write(e.body)
	return err
}
