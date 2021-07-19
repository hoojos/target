package middleware

import "github.com/valyala/fasthttp"

type Middleware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Chain(m ...Middleware) Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}
