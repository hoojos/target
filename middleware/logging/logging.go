package logging

import (
	"time"

	"github.com/hoojos/target/middleware"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func Use(logger logrus.FieldLogger) middleware.Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			start := time.Now()
			next(ctx)
			logger.WithFields(logrus.Fields{
				"remote_addr": ctx.RemoteAddr(),
				"host":        string(ctx.Host()),
				"method":      string(ctx.Method()),
				"path":        string(ctx.Path()),
				"User-Agent":  string(ctx.UserAgent()),
				"latency":     time.Now().Sub(start).Milliseconds(),
				"status":      ctx.Response.StatusCode(),
			}).Info()
		}
	}
}
