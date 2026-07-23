package aliases

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cataliasescolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Aliases struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewAliases func() *Aliases

func NewAliasesFunc(tp elastictransport.Interface) NewAliases {
	_ = "STUB: not implemented"
	return *new(NewAliases)
}

func New(tp elastictransport.Interface) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Aliases) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Aliases) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Aliases) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Aliases) Header(key, value string) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Name(name string) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) H(cataliasescolumns ...cataliasescolumn.CatAliasesColumn) *Aliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *Aliases) S(names ...string) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Aliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *Aliases) MasterTimeout(duration string) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Bytes(bytes bytes.Bytes) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Format(format string) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Help(help bool) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Time(time timeunit.TimeUnit) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) V(v bool) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) ErrorTrace(errortrace bool) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) FilterPath(filterpaths ...string) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Human(human bool) *Aliases { _ = "STUB: not implemented"; return nil }

func (r *Aliases) Pretty(pretty bool) *Aliases { _ = "STUB: not implemented"; return nil }
