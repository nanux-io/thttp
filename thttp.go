package thttp

import (
	"errors"
	"fmt"

	"github.com/nanux-io/nanux"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type httpRoute struct {
	route  string
	method string
}

// Transporter define a tnng instance of transporter which resolve the `Transporter`
// interface from nanux transporter package
type Transporter struct {
	// url on which the http server will listen
	url    string
	Server *fasthttp.Server

	okOptions     bool
	routeHandlers map[httpRoute]nanux.THandler
	errHandler    nanux.ErrorHandler
	closeChan     chan bool
}

// Run start the http server and make it listens on the transporter's url
func (t *Transporter) Run() (err error) {
	t.Server.Handler = func(ctx *fasthttp.RequestCtx) {
		var resp []byte
		var err error
		method := string(ctx.Method())

		log.Debug().Msgf("Receive request for path: %s and method : %s", ctx.Path(), ctx.Method())

		// if option okOptions is set on the transporter then respond 200 to all
		// option request
		if t.okOptions == true && method == fasthttp.MethodOptions {
			ctx.SetStatusCode(200)
			ctx.SetConnectionClose()
			return
		}

		key := httpRoute{
			route:  string(ctx.Path()),
			method: method,
		}

		tHandler, ok := t.routeHandlers[key]

		// if handler not found for path then response with status code 404 is sent
		if ok == false {
			ctx.SetStatusCode(404)
			ctx.SetConnectionClose()
			return
		}

		// create nanux request and provide it with the fasthttp context
		req := nanux.Request{
			Data: ctx.Request.Body(),
			M:    map[string]interface{}{"httpCtx": ctx},
		}

		resp, err = tHandler.Fn(req)

		// in case of error during the execution of the handler, the error handler
		// is called if it is defined, otherwise a 500 status code is set and the
		// response is sent
		if err != nil {
			if t.errHandler == nil {
				ctx.SetStatusCode(500)
				ctx.SetConnectionClose()
				return
			}

			resp = t.errHandler(err, req)

			// if the error handler return a response then the body is set to this value
			// and the status code is set to 500
			if resp != nil {
				ctx.SetStatusCode(500)
				ctx.SetBody(resp)
			}

			ctx.SetConnectionClose()

			return
		}

		// if the handler return a non null value then the body of the response is
		// set with this value
		if resp != nil {
			ctx.SetBody(resp)
		}

		ctx.SetConnectionClose()
		return
	}

	return t.Server.ListenAndServe(t.url)
}

// Close the http server
func (t *Transporter) Close() (err error) {
	if err = t.Server.Shutdown(); err != nil {
		return err
	}

	return
}

// Handle add handler for specified route
func (t *Transporter) Handle(route string, tHandler nanux.THandler) error {
	methodsI, ok := tHandler.Opts[MethodsOpt]

	if ok == false {
		errMsg := fmt.Sprintf("Missing http method for route : %s", route)
		log.Error().Msg(errMsg)

		return errors.New(errMsg)
	}

	methods, ok := methodsI.(Methods)

	if ok == false {
		errMsg := "Option associated to thttp.MethodsOpt is not of type thttp.Methods"
		log.Error().Msg(errMsg)

		return errors.New(errMsg)
	}

	httpRoutes := methods.getHTTPRoutes(route)

	for _, httpRoute := range httpRoutes {
		if _, ok := t.routeHandlers[httpRoute]; ok == true {
			errMsg := "An handler is already associated to this route"
			log.Error().Msg(errMsg)

			return errors.New(errMsg)
		}

		t.routeHandlers[httpRoute] = tHandler
	}

	return nil
}

// HandleError manage error of a handler
func (t *Transporter) HandleError(errHandler nanux.ErrorHandler) (err error) {
	if t.errHandler != nil {
		errMsg := "An error handler has already been set"
		log.Error().Msg(errMsg)

		return errors.New(errMsg)
	}

	t.errHandler = errHandler
	return nil
}

/*----------------------------------------------------------------------------*\
  Instantiation of tHTTP transporter
\*----------------------------------------------------------------------------*/

// New returns a new instance of http transporter which will listen to the specified url.
// The param okOption is a little helper to tell the transporter to respond ok to all
// options.
func New(url string, okOptions bool) Transporter {
	return Transporter{
		url:           url,
		Server:        &fasthttp.Server{},
		routeHandlers: make(map[httpRoute]nanux.THandler),
		okOptions:     okOptions,
	}
}
