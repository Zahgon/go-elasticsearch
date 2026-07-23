package updatecrossclusterapikey

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

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateCrossClusterApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateCrossClusterApiKey func(id string) *UpdateCrossClusterApiKey

func NewUpdateCrossClusterApiKeyFunc(tp elastictransport.Interface) NewUpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return *new(NewUpdateCrossClusterApiKey)
}

func New(tp elastictransport.Interface) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Raw(raw io.Reader) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Request(req *Request) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateCrossClusterApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateCrossClusterApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateCrossClusterApiKey) Header(key, value string) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) _id(id string) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) ErrorTrace(errortrace bool) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) FilterPath(filterpaths ...string) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Human(human bool) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Pretty(pretty bool) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Access(access types.AccessVariant) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) CertificateIdentity(certificateidentity string) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Expiration(duration types.DurationVariant) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateCrossClusterApiKey) Metadata(metadata types.MetadataVariant) *UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}
