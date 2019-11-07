package thttp

import (
	"reflect"
	"testing"

	"github.com/nanux-io/nanux"
	"github.com/valyala/fasthttp"
)

func TestMethods_getHTTPRoutes(t *testing.T) {
	type fields struct {
		Get     bool
		Post    bool
		Put     bool
		Patch   bool
		Delete  bool
		Options bool
		Head    bool
		All     bool
	}
	type args struct {
		route string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantHTTPRoutes []httpRoute
	}{
		{
			name:   "ALL",
			fields: fields{All: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodGet},
				{route: "aroute", method: fasthttp.MethodPost},
				{route: "aroute", method: fasthttp.MethodPut},
				{route: "aroute", method: fasthttp.MethodPatch},
				{route: "aroute", method: fasthttp.MethodDelete},
				{route: "aroute", method: fasthttp.MethodHead},
				{route: "aroute", method: fasthttp.MethodOptions},
			},
		},
		{
			name:   "GET",
			fields: fields{Get: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodGet},
			},
		},
		{
			name:   "POST",
			fields: fields{Post: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodPost},
			},
		},
		{
			name:   "PUT",
			fields: fields{Put: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodPut},
			},
		},
		{
			name:   "PATCH",
			fields: fields{Patch: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodPatch},
			},
		},
		{
			name:   "DELETE",
			fields: fields{Delete: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodDelete},
			},
		},
		{
			name:   "OPTIONS",
			fields: fields{Options: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodOptions},
			},
		},
		{
			name:   "HEAD",
			fields: fields{Head: true},
			args:   args{route: "aroute"},
			wantHTTPRoutes: []httpRoute{
				{route: "aroute", method: fasthttp.MethodHead},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Methods{
				Get:     tt.fields.Get,
				Post:    tt.fields.Post,
				Put:     tt.fields.Put,
				Patch:   tt.fields.Patch,
				Delete:  tt.fields.Delete,
				Options: tt.fields.Options,
				Head:    tt.fields.Head,
				All:     tt.fields.All,
			}
			if gotHTTPRoutes := m.getHTTPRoutes(tt.args.route); !reflect.DeepEqual(gotHTTPRoutes, tt.wantHTTPRoutes) {
				t.Errorf("Methods.getHTTPRoutes() = %v, want %v", gotHTTPRoutes, tt.wantHTTPRoutes)
			}
		})
	}
}

func TestGET(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for get request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Get: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GET(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("GET().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestPOST(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for post request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Post: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := POST(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("Post().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestPut(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for put request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Put: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PUT(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("PUT().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestPATCH(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for patch request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Patch: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PATCH(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("PATCH().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestDELETE(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for delete request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Delete: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DELETE(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("DELETE().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestHEAD(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for head request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Head: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HEAD(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("HEAD().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestOPTIONS(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for options request",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{Options: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OPTIONS(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("OPTIONS().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}

func TestALL(t *testing.T) {
	type args struct {
		fn nanux.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want nanux.Handler
	}{
		{
			name: "return a nanux handler for all request verbs",
			args: args{
				fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
			},
			want: nanux.Handler{
				Fn: func(ctx *interface{}, request nanux.Request) (response []byte, err error) {
					return []byte("custom message"), err
				},
				Opts: nanux.HandlerOpts{MethodsOpt: Methods{All: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ALL(tt.args.fn)

			resp, err := got.Fn(nil, nanux.Request{})

			if string(resp) != "custom message" || err != nil {
				t.Errorf("Fn argument has not be successfully provided to nanux.Handler")
			}

			if !reflect.DeepEqual(got.Opts, tt.want.Opts) {
				t.Errorf("ALL().Opts = %v, want %v", got.Opts, tt.want.Opts)
			}
		})
	}
}
