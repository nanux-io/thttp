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
		wantErr     bool
	}{
		{
			name:    "http ctx not provided",
			args:    args{req: nanux.Request{M: make(map[string]interface{})}},
			wantErr: true,
		},
		{
			name:    "http ctx is not of type *RequestCtx",
			args:    args{req: nanux.Request{M: map[string]interface{}{"httpCtx": "wrong type"}}},
			wantErr: true,
		},
		{
			name:        "http ctx type is *RequestCtx",
			args:        args{req: nanux.Request{M: map[string]interface{}{"httpCtx": &fasthttp.RequestCtx{}}}},
			wantHTTPCtx: &fasthttp.RequestCtx{},
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHTTPCtx, gotErr := GetHTTPCtx(tt.args.req)
			if !reflect.DeepEqual(gotHTTPCtx, tt.wantHTTPCtx) {
				t.Errorf("GetHTTPCtx() gotHttpCtx = %v, want %v", gotHTTPCtx, tt.wantHTTPCtx)
			}
			if (tt.wantErr == true && gotErr == nil) || (tt.wantErr == false && gotErr != nil) {
				t.Errorf("GetHTTPCtx() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
