package migratereindex

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/modeenum"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MigrateReindex struct {
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

type NewMigrateReindex func() *MigrateReindex

func NewMigrateReindexFunc(tp elastictransport.Interface) NewMigrateReindex {
	_ = "STUB: not implemented"
	return *new(NewMigrateReindex)
}

func New(tp elastictransport.Interface) *MigrateReindex { _ = "STUB: not implemented"; return nil }

func (r *MigrateReindex) Raw(raw io.Reader) *MigrateReindex { _ = "STUB: not implemented"; return nil }

func (r *MigrateReindex) Request(req *Request) *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateReindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateReindex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateReindex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *MigrateReindex) Header(key, value string) *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateReindex) ErrorTrace(errortrace bool) *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateReindex) FilterPath(filterpaths ...string) *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateReindex) Human(human bool) *MigrateReindex { _ = "STUB: not implemented"; return nil }

func (r *MigrateReindex) Pretty(pretty bool) *MigrateReindex { _ = "STUB: not implemented"; return nil }

func (r *MigrateReindex) Mode(mode modeenum.ModeEnum) *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateReindex) Source(source types.SourceIndexVariant) *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}
