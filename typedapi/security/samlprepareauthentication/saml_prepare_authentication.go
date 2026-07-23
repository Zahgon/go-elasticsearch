package samlprepareauthentication

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SamlPrepareAuthentication struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSamlPrepareAuthentication func() *SamlPrepareAuthentication

func NewSamlPrepareAuthenticationFunc(tp elastictransport.Interface) NewSamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return *new(NewSamlPrepareAuthentication)
}

func New(tp elastictransport.Interface) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) Raw(raw io.Reader) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) Request(req *Request) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlPrepareAuthentication) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlPrepareAuthentication) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SamlPrepareAuthentication) Header(key, value string) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) ErrorTrace(errortrace bool) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) FilterPath(filterpaths ...string) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) Human(human bool) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) Pretty(pretty bool) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) Acs(acs string) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) Realm(realm string) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlPrepareAuthentication) RelayState(relaystate string) *SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}
