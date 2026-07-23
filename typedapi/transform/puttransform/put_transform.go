package puttransform

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

type PutTransform struct {
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

type NewPutTransform func(transformid string) *PutTransform

func NewPutTransformFunc(tp elastictransport.Interface) NewPutTransform {
	_ = "STUB: not implemented"
	return *new(NewPutTransform)
}

func New(tp elastictransport.Interface) *PutTransform { _ = "STUB: not implemented"; return nil }

func (r *PutTransform) Raw(raw io.Reader) *PutTransform { _ = "STUB: not implemented"; return nil }

func (r *PutTransform) Request(req *Request) *PutTransform { _ = "STUB: not implemented"; return nil }

func (r *PutTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutTransform) Header(key, value string) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) _transformid(transformid string) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) DeferValidation(defervalidation bool) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Timeout(duration string) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) ErrorTrace(errortrace bool) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) FilterPath(filterpaths ...string) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Human(human bool) *PutTransform { _ = "STUB: not implemented"; return nil }

func (r *PutTransform) Pretty(pretty bool) *PutTransform { _ = "STUB: not implemented"; return nil }

func (r *PutTransform) Description(description string) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Dest(dest types.TransformDestinationVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Frequency(duration types.DurationVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Latest(latest types.LatestVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Meta_(metadata types.MetadataVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Pivot(pivot types.PivotVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) RetentionPolicy(retentionpolicy types.RetentionPolicyContainerVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Settings(settings types.SettingsVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Source(source types.TransformSourceVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTransform) Sync(sync types.SyncContainerVariant) *PutTransform {
	_ = "STUB: not implemented"
	return nil
}
