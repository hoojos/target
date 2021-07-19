package endpoint

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Endpoint struct {
	method  string
	URL     string
	headers map[string]string
	body    []byte
	latency time.Duration
	metric  Metrics
}

func (e Endpoint) Handle(ctx *fasthttp.RequestCtx) {
	time.Sleep(e.latency)
	for k, v := range e.headers {
		ctx.Response.Header.Set(k, v)
	}
	_, _ = ctx.Write(e.body)
}
