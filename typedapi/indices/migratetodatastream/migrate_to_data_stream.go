package migratetodatastream

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MigrateToDataStream struct {
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

type NewMigrateToDataStream func(name string) *MigrateToDataStream

func NewMigrateToDataStreamFunc(tp elastictransport.Interface) NewMigrateToDataStream {
	_ = "STUB: not implemented"
	return *new(NewMigrateToDataStream)
}

func New(tp elastictransport.Interface) *MigrateToDataStream { _ = "STUB: not implemented"; return nil }

func (r *MigrateToDataStream) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateToDataStream) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateToDataStream) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateToDataStream) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *MigrateToDataStream) Header(key, value string) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) _name(name string) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) MasterTimeout(duration string) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) Timeout(duration string) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) ErrorTrace(errortrace bool) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) FilterPath(filterpaths ...string) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) Human(human bool) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataStream) Pretty(pretty bool) *MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}
