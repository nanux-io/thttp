package thttp

import (
	"reflect"
	"testing"

	"github.com/nanux-io/nanux"
	"github.com/valyala/fasthttp"
)

func TestGetHTTPCtx(t *testing.T) {
	type args struct {
		req nanux.Request
	}

	tests := []struct {
		name        string
		args        args
		wantHTTPCtx *fasthttp.RequestCtx
		wantOk      bool
	}{
		{
			name:   "http ctx not provided",
			args:   args{req: nanux.Request{M: make(map[string]interface{})}},
			wantOk: false,
		},
		{
			name:   "http ctx is not of type *RequestCtx",
			args:   args{req: nanux.Request{M: map[string]interface{}{"httpCtx": "wrong type"}}},
			wantOk: false,
		},
		{
			name:        "http ctx type is *RequestCtx",
			args:        args{req: nanux.Request{M: map[string]interface{}{"httpCtx": &fasthttp.RequestCtx{}}}},
			wantHTTPCtx: &fasthttp.RequestCtx{},
			wantOk:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHTTPCtx, gotOk := GetHTTPCtx(tt.args.req)
			if !reflect.DeepEqual(gotHTTPCtx, tt.wantHTTPCtx) {
				t.Errorf("GetHTTPCtx() gotHttpCtx = %v, want %v", gotHTTPCtx, tt.wantHTTPCtx)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetHTTPCtx() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
