package createcrossclusterapikey

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateCrossClusterApiKey struct {
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

type NewCreateCrossClusterApiKey func() *CreateCrossClusterApiKey

func NewCreateCrossClusterApiKeyFunc(tp elastictransport.Interface) NewCreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return *new(NewCreateCrossClusterApiKey)
}

func New(tp elastictransport.Interface) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Raw(raw io.Reader) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Request(req *Request) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateCrossClusterApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateCrossClusterApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CreateCrossClusterApiKey) Header(key, value string) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) ErrorTrace(errortrace bool) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) FilterPath(filterpaths ...string) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Human(human bool) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Pretty(pretty bool) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Access(access types.AccessVariant) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) CertificateIdentity(certificateidentity string) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Expiration(duration types.DurationVariant) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Metadata(metadata types.MetadataVariant) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateCrossClusterApiKey) Name(name string) *CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}
