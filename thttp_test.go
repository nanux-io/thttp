package thttp_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"

	"github.com/nanux-io/nanux"
	. "github.com/nanux-io/thttp"
)

var _ = Describe("tHTTP transporter", func() {
	httpClient := http.Client{Timeout: 100 * time.Millisecond}
	url := "127.0.0.1:1234"

	It("should create a new instance", func() {
		t := New(url, false)

		var i interface{} = &t
		_, ok := i.(nanux.Transporter)
		Expect(ok).To(Equal(true))
	})

	Context("instance", func() {
		var (
			t         Transporter
			okOptions bool
		)

		JustBeforeEach(func() {
			t = New(url, okOptions)
		})

		It("should launch an http server on the specified url and close", func(done Done) {
			go t.Run()

			// wait to let time for the http server to be launched
			time.Sleep(50 * time.Millisecond)

			resp, err := httpClient.Get("http://" + url)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(404))

			t.Close()

			close(done)
		}, 0.5)

		It("should close a running instance", func() {
			var err error

			go t.Run()
			// wait to let time for the http server to be launched
			time.Sleep(50 * time.Millisecond)

			err = t.Close()
			Expect(err).ToNot(HaveOccurred())

			_, err = httpClient.Get("http://" + url)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Get http://" + url + ": dial tcp " + url + ": connect: connection refused"))

		})

		Context("running", func() {
			methodGetOpt := nanux.HandlerOpts{MethodsOpt: Methods{Get: true}}

			JustBeforeEach(func() {
				go t.Run()

				// wait to let time to the server to be launcher
				time.Sleep(50 * time.Millisecond)
			})

			Context("with ok options", func() {
				BeforeEach(func() {
					okOptions = true
				})

				It("should respond with 200 for all OPTIONS requests", func() {
					req, err := http.NewRequest(http.MethodOptions, "http://"+url, nil)
					Expect(err).ToNot(HaveOccurred())

					res, err := httpClient.Do(req)
					Expect(err).ToNot(HaveOccurred())
					Expect(res.StatusCode).To(Equal(200))

					body, err := readResponseBody(res)
					Expect(err).ToNot(HaveOccurred())
					Expect(body).To(Equal(""))
				})
			})

			It("should raise an error if try to run an already running instance", func(done Done) {
				errC := make(chan error)

				go func() {
					errC <- t.Run()
				}()

				Expect(<-errC).To(HaveOccurred())
				close(done)
			}, 0.2)

			Context("adding handler", func() {
				route := "/allowaddHandler/route"
				tHandler := nanux.THandler{
					Fn: func(nanux.Request) ([]byte, error) {
						return nil, nil
					},
					Opts: methodGetOpt,
				}

				It("should be added the first time and rising an error the second time for the same route", func() {
					var err error

					err = t.Handle(route, tHandler)
					Expect(err).ToNot(HaveOccurred())

					err = t.Handle(route, tHandler)
					Expect(err).To(HaveOccurred())
				})

				It("should fail if the method option is not set in the options of the handler", func() {
					tHandlerWithoutOpt := nanux.THandler{
						Fn: func(nanux.Request) ([]byte, error) {
							return nil, nil
						},
						Opts: nanux.HandlerOpts{MethodsOpt: "wrong type"},
					}

					err := t.Handle("routewithoutopt", tHandlerWithoutOpt)
					Expect(err).To(HaveOccurred())
				})

				It("should fail if the method option is wrong type in the options of the handler", func() {
					tHandlerWithoutOpt := nanux.THandler{
						Fn: func(nanux.Request) ([]byte, error) {
							return nil, nil
						},
					}

					err := t.Handle("routewithoutopt", tHandlerWithoutOpt)
					Expect(err).To(HaveOccurred())
				})
			})

			It("should inject fasthttp context into Request.M of handler", func(done Done) {
				route := "/test/req"
				c := make(chan bool)
				tHandler := nanux.THandler{
					Fn: func(req nanux.Request) ([]byte, error) {
						var ok bool
						var httpCtx interface{}

						httpCtx, ok = req.M["httpCtx"]

						if ok == false {
							c <- false
							return nil, nil
						}

						_, ok = httpCtx.(*fasthttp.RequestCtx)

						if ok == false {
							c <- false
						}

						c <- true

						return nil, nil
					},
					Opts: methodGetOpt,
				}

				err := t.Handle(route, tHandler)
				Expect(err).ToNot(HaveOccurred())

				go httpClient.Get("http://" + url + route)
				Expect(<-c).To(Equal(true))

				close(done)
			}, 0.5)

			It("should set into the response body the value returned by the handler", func() {
				handlerMsg := "message coming from my handler"
				route := "/myroute"
				tHandler := nanux.THandler{
					Fn: func(req nanux.Request) ([]byte, error) {
						return []byte(handlerMsg), nil
					},
					Opts: methodGetOpt,
				}

				err := t.Handle(route, tHandler)
				Expect(err).ToNot(HaveOccurred())

				resp, err := httpClient.Get("http://" + url + route)
				Expect(err).ToNot(HaveOccurred())

				body, _ := readResponseBody(resp)
				Expect(resp.StatusCode).To(Equal(200))
				Expect(body).To(Equal(handlerMsg))
			})

			It("should only respond to methods (GET, DELETE) set into the options of the handler", func() {
				route := "/myroute"
				tHandler := nanux.THandler{
					Fn: func(req nanux.Request) ([]byte, error) {
						return nil, nil
					},
					Opts: nanux.HandlerOpts{MethodsOpt: Methods{Get: true, Delete: true}},
				}

				err := t.Handle(route, tHandler)
				Expect(err).ToNot(HaveOccurred())

				resp, err := httpClient.Get("http://" + url + route)
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))

				req, err := http.NewRequest(http.MethodDelete, "http://"+url+route, nil)
				Expect(err).ToNot(HaveOccurred())
				resp, err = httpClient.Do(req)

				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))
			})

			It("should allow to add errorHandler only once", func() {
				var err error

				errHandler := func(error, nanux.Request) []byte {
					return nil
				}

				err = t.HandleError(errHandler)
				Expect(err).ToNot(HaveOccurred())
				err = t.HandleError(errHandler)
				Expect(err).To(HaveOccurred())
			})

			Context("when the handler raise an error", func() {
				var handlerErrMsg string
				route := "/test/route"
				routeFullUrl := "http://" + url + route

				JustBeforeEach(func() {
					handlerErrMsg = "fake error in request handler"
					tHandler := nanux.THandler{
						Fn: func(nanux.Request) ([]byte, error) {
							return nil, errors.New(handlerErrMsg)
						},
						Opts: methodGetOpt,
					}

					err := t.Handle(route, tHandler)
					Expect(err).ToNot(HaveOccurred())
				})

				It("should respond with 500 status when there is no error handler", func() {
					resp, err := httpClient.Get(routeFullUrl)
					Expect(err).ToNot(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(500))

					body, err := readResponseBody(resp)
					Expect(body).To(Equal(""))
				})

				Context("when the error handler return a value", func() {
					It("should respond with the value provided by the error handler and with 500 status code", func() {
						errHandler := func(err error, req nanux.Request) []byte {
							return []byte(err.Error())
						}

						err := t.HandleError(errHandler)
						Expect(err).ToNot(HaveOccurred())

						resp, err := httpClient.Get(routeFullUrl)
						Expect(err).ToNot(HaveOccurred())
						Expect(resp.StatusCode).To(Equal(500))

						body, err := readResponseBody(resp)
						Expect(body).To(Equal(handlerErrMsg))
					})
				})

				Context("when the error handler does not return a value", func() {
					It("should just close the connection", func() {
						errMsg := "err message set into the response body from the error handler"
						errHandler := func(err error, req nanux.Request) []byte {
							httpCtxI, ok := req.M["httpCtx"]

							if ok == false {
								return []byte("http context not present in req.M")
							}

							httpCtx, ok := httpCtxI.(*fasthttp.RequestCtx)

							if ok == false {
								return []byte("can not cast http context")
							}

							httpCtx.SetStatusCode(400)
							httpCtx.SetBody([]byte(errMsg))

							return nil
						}

						err := t.HandleError(errHandler)
						Expect(err).ToNot(HaveOccurred())

						resp, err := httpClient.Get(routeFullUrl)
						Expect(err).ToNot(HaveOccurred())
						Expect(resp.StatusCode).To(Equal(400))

						body, _ := readResponseBody(resp)
						Expect(body).To(Equal(errMsg))
					})
				})
			})

			AfterEach(func() {
				t.Close()
			})
		})
	})
})

func readResponseBody(resp *http.Response) (body string, err error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	body = string(bodyBytes)

	return
}
