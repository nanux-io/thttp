# THTTP

tHTTP is a http transporter to be used with [nanux](https://github.com/nanux-io/nanux)

tHTTP is based on [fasthttp](https://github.com/valyala/fasthttp) to provide the http server. fasthttp was chosen and not the std lib because fasthttp is faster (see bench).

code coverage: 97.7%

## Usage

### Creation

To create a new http transporter the `New` method must be called. It takes 2 
arguments:

* **url** which define on which url the http server will listen
* **okOptions** which if set to true tell the server to respond with 200 status code 
to all OPTIONS requests.

```go

func creatingHTTPTransporter() nanux.Transporter {
  return thttp.New("127.0.0.1:8000", true)
}

```

### Handlers

tHTTP inject the instant of `*fasthttp.RequestCtx` in `req.M["httpCtx"]` where 
`req` is the nanux request of the nanux.HandlerFunc. To get it as an instance of
`*fasthttp.RequestCtx` there is a helper: `thttp.GetHTTPCtx(req)`

```go
handler := nanux.Handler{
  Fn: func(ctx *interface{}, req Request) ([]byte, error) {
    // httpCtxI is an interface corresponding to a `*fasthttp.RequestCtx`
    httpCtxI := req.M["httpCtx"] 

    // httpCtx is an instance of `*fasthttp.RequestCtx` if `ok` is `true`
    // otherwise it is nil
    httpCtx, ok := thttp.GetHTTPCtx(req)
  }
}
```

### Middlewares

Official middlewares:

* **OKOptions**: make a default response to Options request. If it is used
in combination with a `EnsureMETHOD` middleware, be sure to call `OKOptions` first

## Development

Command to execute test: `go test -coverprofile=coverage.out -v &&  go tool cover -html=coverage.out -o coverage.html`  
Command to execute test with check race condition: `go test -race -coverprofile=coverage.out -v &&  go tool cover -html=coverage.out -o coverage.html`

## Contributor

Thanks to [Nicolas Talle](https://github.com/nicolab) for the feedback.
