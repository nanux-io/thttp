# THTTP

tHTTP is a http transporter to be used with [nanux](https://github.com/nanux-io/nanux)

tHTTP is based on [fasthttp](https://github.com/valyala/fasthttp) to provide the http server.

code coverage: 97.7%

## Usage

### Creation

```go

func creatingHTTPTransporter() nanux.Transporter {
  return thttp.New("127.0.0.1:8000")
}

```

### Actions

```go
req.M["httpCtx"] // fatshttp context
```

## Development

Command to execute test: `go test -coverprofile=coverage.out -v &&  go tool cover -html=coverage.out -o coverage.html`
Command to execute test with check race condition: `go test -race -coverprofile=coverage.out -v &&  go tool cover -html=coverage.out -o coverage.html`

## Contributor

Thanks to [Nicolas Talle](https://github.com/nicolab) for the feedback.
