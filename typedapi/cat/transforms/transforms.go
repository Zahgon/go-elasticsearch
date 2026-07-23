package transforms

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cattransformcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	transformidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Transforms struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTransforms func() *Transforms

func NewTransformsFunc(tp elastictransport.Interface) NewTransforms {
	_ = "STUB: not implemented"
	return *new(NewTransforms)
}

func New(tp elastictransport.Interface) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Transforms) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Transforms) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Transforms) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Transforms) Header(key, value string) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) TransformId(transformid string) *Transforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *Transforms) AllowNoMatch(allownomatch bool) *Transforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *Transforms) From(from int) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) H(cattransformcolumns ...cattransformcolumn.CatTransformColumn) *Transforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *Transforms) S(cattransformcolumns ...cattransformcolumn.CatTransformColumn) *Transforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *Transforms) Size(size int) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) Bytes(bytes bytes.Bytes) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) Format(format string) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) Help(help bool) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) Time(time timeunit.TimeUnit) *Transforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *Transforms) V(v bool) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) ErrorTrace(errortrace bool) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) FilterPath(filterpaths ...string) *Transforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *Transforms) Human(human bool) *Transforms { _ = "STUB: not implemented"; return nil }

func (r *Transforms) Pretty(pretty bool) *Transforms { _ = "STUB: not implemented"; return nil }
