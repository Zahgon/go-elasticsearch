package previewtransform

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

type PreviewTransform struct {
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

type NewPreviewTransform func() *PreviewTransform

func NewPreviewTransformFunc(tp elastictransport.Interface) NewPreviewTransform {
	_ = "STUB: not implemented"
	return *new(NewPreviewTransform)
}

func New(tp elastictransport.Interface) *PreviewTransform { _ = "STUB: not implemented"; return nil }

func (r *PreviewTransform) Raw(raw io.Reader) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Request(req *Request) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PreviewTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PreviewTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PreviewTransform) Header(key, value string) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) TransformId(transformid string) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Timeout(duration string) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) ErrorTrace(errortrace bool) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) FilterPath(filterpaths ...string) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Human(human bool) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Pretty(pretty bool) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Description(description string) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Dest(dest types.TransformDestinationVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Frequency(duration types.DurationVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Latest(latest types.LatestVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Pivot(pivot types.PivotVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) RetentionPolicy(retentionpolicy types.RetentionPolicyContainerVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Settings(settings types.SettingsVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Source(source types.TransformSourceVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewTransform) Sync(sync types.SyncContainerVariant) *PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}
