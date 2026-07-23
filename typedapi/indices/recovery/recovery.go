package recovery

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Recovery struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRecovery func() *Recovery

func NewRecoveryFunc(tp elastictransport.Interface) NewRecovery {
	_ = "STUB: not implemented"
	return *new(NewRecovery)
}

func New(tp elastictransport.Interface) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Recovery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Recovery) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Recovery) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Recovery) Header(key, value string) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) Index(index string) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) ActiveOnly(activeonly bool) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) Detailed(detailed bool) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) AllowNoIndices(allownoindices bool) *Recovery {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recovery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Recovery {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recovery) IgnoreUnavailable(ignoreunavailable bool) *Recovery {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recovery) ErrorTrace(errortrace bool) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) FilterPath(filterpaths ...string) *Recovery {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recovery) Human(human bool) *Recovery { _ = "STUB: not implemented"; return nil }

func (r *Recovery) Pretty(pretty bool) *Recovery { _ = "STUB: not implemented"; return nil }
