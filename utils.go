package thttp

import (
	"github.com/nanux-io/nanux"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

// GetHTTPCtx return the fasthttp request context extract from the nanux request
func GetHTTPCtx(req nanux.Request) (httpCtx *fasthttp.RequestCtx, ok bool) {
	httpCtxI, ok := req.M["httpCtx"]

	if ok == false {
		log.Error().Msg("GetHTTPCtx : could not extract http context from request")
		return
	}

	httpCtx, ok = httpCtxI.(*fasthttp.RequestCtx)

	if ok == false {
		log.Error().Msg("GetHTTPCtx : could not convert http context to *fasthttp.RequestCtx")

		return
	}

	return
}
