package thttp

import (
	"testing"

	"github.com/nanux-io/nanux"
	"github.com/valyala/fasthttp"
)

func TestOKOptions(t *testing.T) {
	type args struct {
		fn         nanux.HandlerFunc
		getHTTPCtx func() *fasthttp.RequestCtx
	}
	tests := []struct {
		name           string
		args           args
		wantErr        error
		wantRes        []byte
		wantStatusCode int
	}{
		{
			name: "request option",
			args: args{
				fn: func(*interface{}, nanux.Request) ([]byte, error) {
					return []byte("some response"), nil
				},
				getHTTPCtx: func() *fasthttp.RequestCtx {
					h := fasthttp.RequestHeader{}
					h.SetMethod("OPTIONS")

					return &fasthttp.RequestCtx{
						Request: fasthttp.Request{
							Header: h,
						},
					}
				},
			},
			wantErr:        nil,
			wantRes:        nil,
			wantStatusCode: 200,
		},
		{
			name: "request get",
			args: args{
				fn: func(*interface{}, nanux.Request) ([]byte, error) {
					return []byte("some response"), nil
				},
				getHTTPCtx: func() *fasthttp.RequestCtx {
					h := fasthttp.RequestHeader{}
					h.SetMethod("GET")

					return &fasthttp.RequestCtx{
						Request: fasthttp.Request{
							Header: h,
						},
					}
				},
			},
			wantErr:        nil,
			wantRes:        []byte("some response"),
			wantStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpCtx := tt.args.getHTTPCtx()
			req := nanux.Request{
				M: map[string]interface{}{"httpCtx": httpCtx},
			}

			got := OKOptions(tt.args.fn)

			res, err := got(nil, req)

			if err != tt.wantErr {
				t.Errorf("OKOptions() - error occured when calling handler - %s", err)
			}

			if string(res) != string(tt.wantRes) {
				t.Error("OKOptions() - return a response when calling handler instead of nil")
			}

			if httpCtx.Response.StatusCode() != tt.wantStatusCode {
				t.Error("OKOptions() - must set status code to 200 for OPTIONS request")
			}
		})
	}
}
