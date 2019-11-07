package thttp

import (
	"github.com/nanux-io/nanux"
	"github.com/valyala/fasthttp"
)

const (
	// MethodsOpt define the handler option key for specifying http method associated
	// to the handler. The value associated to this option come from the http
	// package of the std lib (eg http.MethodGet)
	MethodsOpt nanux.HandlerOptName = "httpMethod"
)

// Methods define available methods for handler
type Methods struct {
	Get     bool
	Post    bool
	Put     bool
	Patch   bool
	Delete  bool
	Options bool
	Head    bool
	All     bool
}

func (m Methods) getHTTPRoutes(route string) (httpRoutes []httpRoute) {
	if m.All == true {
		m.Get = true
		m.Post = true
		m.Put = true
		m.Patch = true
		m.Delete = true
		m.Options = true
		m.Head = true
	}

	if m.Get == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodGet})
	}

	if m.Post == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodPost})
	}

	if m.Put == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodPut})
	}

	if m.Patch == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodPatch})
	}

	if m.Delete == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodDelete})
	}

	if m.Head == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodHead})
	}

	if m.Options == true {
		httpRoutes = append(httpRoutes, httpRoute{route: route, method: fasthttp.MethodOptions})
	}

	return
}

// GET return a handler for the specified handle func which will respond to GET request
func GET(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Get: true}},
	}
}

// POST return a handler for the specified handle func which will respond to POST request
func POST(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Post: true}},
	}
}

// PUT return a handler for the specified handle func which will respond to Put request
func PUT(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Put: true}},
	}
}

// PATCH return a handler for the specified handle func which will respond to PATCH request
func PATCH(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Patch: true}},
	}
}

// DELETE return a handler for the specified handle func which will respond to DELETE request
func DELETE(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Delete: true}},
	}
}

// HEAD return a handler for the specified handle func which will respond to HEAD request
func HEAD(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Head: true}},
	}
}

// OPTIONS return a handler for the specified handle func which will respond to OPTIONS request
func OPTIONS(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{Options: true}},
	}
}

// ALL return a handler for the specified handle func which will respond to all request verb
func ALL(fn nanux.HandlerFunc) nanux.Handler {
	return nanux.Handler{
		Fn:   fn,
		Opts: nanux.HandlerOpts{MethodsOpt: Methods{All: true}},
	}
}
