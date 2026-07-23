package logsdisable

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

type LogsDisable struct {
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

type NewLogsDisable func(name string) *LogsDisable

func NewLogsDisableFunc(tp elastictransport.Interface) NewLogsDisable {
	_ = "STUB: not implemented"
	return *new(NewLogsDisable)
}

func New(tp elastictransport.Interface) *LogsDisable { _ = "STUB: not implemented"; return nil }

func (r *LogsDisable) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LogsDisable) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LogsDisable) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LogsDisable) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *LogsDisable) Header(key, value string) *LogsDisable { _ = "STUB: not implemented"; return nil }

func (r *LogsDisable) _name(name string) *LogsDisable { _ = "STUB: not implemented"; return nil }

func (r *LogsDisable) MasterTimeout(duration string) *LogsDisable {
	_ = "STUB: not implemented"
	return nil
}

func (r *LogsDisable) Timeout(duration string) *LogsDisable { _ = "STUB: not implemented"; return nil }

func (r *LogsDisable) ErrorTrace(errortrace bool) *LogsDisable {
	_ = "STUB: not implemented"
	return nil
}

func (r *LogsDisable) FilterPath(filterpaths ...string) *LogsDisable {
	_ = "STUB: not implemented"
	return nil
}

func (r *LogsDisable) Human(human bool) *LogsDisable { _ = "STUB: not implemented"; return nil }

func (r *LogsDisable) Pretty(pretty bool) *LogsDisable { _ = "STUB: not implemented"; return nil }
