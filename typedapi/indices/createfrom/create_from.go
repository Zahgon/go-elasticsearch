package createfrom

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
	sourceMask = iota + 1

	destMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateFrom struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	source string
	dest   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCreateFrom func(source, dest string) *CreateFrom

func NewCreateFromFunc(tp elastictransport.Interface) NewCreateFrom {
	_ = "STUB: not implemented"
	return *new(NewCreateFrom)
}

func New(tp elastictransport.Interface) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) Raw(raw io.Reader) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) Request(req *Request) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateFrom) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateFrom) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CreateFrom) Header(key, value string) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) _source(source string) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) _dest(dest string) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) ErrorTrace(errortrace bool) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) FilterPath(filterpaths ...string) *CreateFrom {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateFrom) Human(human bool) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) Pretty(pretty bool) *CreateFrom { _ = "STUB: not implemented"; return nil }

func (r *CreateFrom) MappingsOverride(mappingsoverride types.TypeMappingVariant) *CreateFrom {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateFrom) RemoveIndexBlocks(removeindexblocks bool) *CreateFrom {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateFrom) SettingsOverride(settingsoverride types.IndexSettingsVariant) *CreateFrom {
	_ = "STUB: not implemented"
	return nil
}
