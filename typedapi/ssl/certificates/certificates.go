package certificates

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Certificates struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCertificates func() *Certificates

func NewCertificatesFunc(tp elastictransport.Interface) NewCertificates {
	_ = "STUB: not implemented"
	return *new(NewCertificates)
}

func New(tp elastictransport.Interface) *Certificates { _ = "STUB: not implemented"; return nil }

func (r *Certificates) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Certificates) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Certificates) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Certificates) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Certificates) Header(key, value string) *Certificates {
	_ = "STUB: not implemented"
	return nil
}

func (r *Certificates) ErrorTrace(errortrace bool) *Certificates {
	_ = "STUB: not implemented"
	return nil
}

func (r *Certificates) FilterPath(filterpaths ...string) *Certificates {
	_ = "STUB: not implemented"
	return nil
}

func (r *Certificates) Human(human bool) *Certificates { _ = "STUB: not implemented"; return nil }

func (r *Certificates) Pretty(pretty bool) *Certificates { _ = "STUB: not implemented"; return nil }
