package thttp

import (
	"errors"

	"github.com/nanux-io/nanux"
)

// OKOptions respond to the request with an empty body and status code 200
// to options request. Because several libs (in different language) make options
// request before doing the "real" request, this middleware is here to help
// answering these requests
func OKOptions(fn nanux.HandlerFunc) nanux.HandlerFunc {
	return func(ctx *interface{}, req nanux.Request) ([]byte, error) {
		httpCtx, ok := GetHTTPCtx(req)

		if ok == false {
			return nil, errors.New("Internal server error")
		}

		if httpCtx.IsOptions() == true {
			return nil, nil
		}

		return fn(ctx, req)
	}
}