package app

import (
	"github.com/valyala/fasthttp"
	"time"
)

// LogRequests - middleware that logs all requests
func (a *App) LogRequests(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		h(ctx)
		fullURI := string(ctx.Request.URI().Path())
		a.logger.Printf("start request: %s | status code: %d | request duration: %s | ip: %s | method: %s, URI: %s",
			start.Format("2006-01-02 - 15:04:05"), ctx.Response.StatusCode(),
			time.Since(start), ctx.RemoteIP(), ctx.Method(), fullURI)
	}
}
