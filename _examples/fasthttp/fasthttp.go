package fasthttp

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

type Transport struct{}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) copyRequest(dst *fasthttp.Request, src *http.Request) *fasthttp.Request {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) copyResponse(dst *http.Response, src *fasthttp.Response) *http.Response {
	_ = "STUB: not implemented"
	return nil
}
