package updatetransform

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
	transformidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateTransform func(transformid string) *UpdateTransform

func NewUpdateTransformFunc(tp elastictransport.Interface) NewUpdateTransform {
	_ = "STUB: not implemented"
	return *new(NewUpdateTransform)
}

func New(tp elastictransport.Interface) *UpdateTransform { _ = "STUB: not implemented"; return nil }

func (r *UpdateTransform) Raw(raw io.Reader) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Request(req *Request) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateTransform) Header(key, value string) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) _transformid(transformid string) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) DeferValidation(defervalidation bool) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Timeout(duration string) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) ErrorTrace(errortrace bool) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) FilterPath(filterpaths ...string) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Human(human bool) *UpdateTransform { _ = "STUB: not implemented"; return nil }

func (r *UpdateTransform) Pretty(pretty bool) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Description(description string) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Dest(dest types.TransformDestinationVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Frequency(duration types.DurationVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Meta_(metadata types.MetadataVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) RetentionPolicy(retentionpolicy types.RetentionPolicyContainerVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Settings(settings types.SettingsVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Source(source types.TransformSourceVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTransform) Sync(sync types.SyncContainerVariant) *UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}
