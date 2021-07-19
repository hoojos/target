package recovery

import (
	"fmt"
	"github.com/hoojos/target/log"
	"net/http"
	"runtime"

	"github.com/hoojos/target/middleware"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Option func(*options)

type options struct {
	handler fasthttp.RequestHandler
	logger  logrus.FieldLogger
}

func WithLogger(logger logrus.FieldLogger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func WithHandler(handler fasthttp.RequestHandler) Option {
	return func(o *options) {
		o.handler = handler
	}
}

func Recovery(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
		handler: func(ctx *fasthttp.RequestCtx) {
			ctx.Error(fmt.Sprintf("panic triggered"), http.StatusInternalServerError)
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			defer func() {
				if p := recover(); p != nil {
					buf := make([]byte, 1024)
					n := runtime.Stack(buf, false)
					buf = buf[:n]
					options.logger.WithFields(logrus.Fields{"panic": p, "stack": string(buf)}).Error("panic")
					options.handler(ctx)
				}
			}()
			next(ctx)
		}
	}
}
