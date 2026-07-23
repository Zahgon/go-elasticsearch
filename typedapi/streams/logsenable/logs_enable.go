package logsenable

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

type LogsEnable struct {
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

type NewLogsEnable func(name string) *LogsEnable

func NewLogsEnableFunc(tp elastictransport.Interface) NewLogsEnable {
	_ = "STUB: not implemented"
	return *new(NewLogsEnable)
}

func New(tp elastictransport.Interface) *LogsEnable { _ = "STUB: not implemented"; return nil }

func (r *LogsEnable) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LogsEnable) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LogsEnable) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LogsEnable) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *LogsEnable) Header(key, value string) *LogsEnable { _ = "STUB: not implemented"; return nil }

func (r *LogsEnable) _name(name string) *LogsEnable { _ = "STUB: not implemented"; return nil }

func (r *LogsEnable) MasterTimeout(duration string) *LogsEnable {
	_ = "STUB: not implemented"
	return nil
}

func (r *LogsEnable) Timeout(duration string) *LogsEnable { _ = "STUB: not implemented"; return nil }

func (r *LogsEnable) ErrorTrace(errortrace bool) *LogsEnable { _ = "STUB: not implemented"; return nil }

func (r *LogsEnable) FilterPath(filterpaths ...string) *LogsEnable {
	_ = "STUB: not implemented"
	return nil
}

func (r *LogsEnable) Human(human bool) *LogsEnable { _ = "STUB: not implemented"; return nil }

func (r *LogsEnable) Pretty(pretty bool) *LogsEnable { _ = "STUB: not implemented"; return nil }
