package executewatch

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actionexecutionmode"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExecuteWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewExecuteWatch func() *ExecuteWatch

func NewExecuteWatchFunc(tp elastictransport.Interface) NewExecuteWatch {
	_ = "STUB: not implemented"
	return *new(NewExecuteWatch)
}

func New(tp elastictransport.Interface) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) Raw(raw io.Reader) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) Request(req *Request) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ExecuteWatch) Header(key, value string) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) Id(id string) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) Debug(debug bool) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) ErrorTrace(errortrace bool) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) FilterPath(filterpaths ...string) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) Human(human bool) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) Pretty(pretty bool) *ExecuteWatch { _ = "STUB: not implemented"; return nil }

func (r *ExecuteWatch) ActionModes(actionmodes map[string]actionexecutionmode.ActionExecutionMode) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) AddActionMode(key string, value actionexecutionmode.ActionExecutionMode) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) AlternativeInput(alternativeinput map[string]json.RawMessage) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) AddAlternativeInput(key string, value json.RawMessage) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) IgnoreCondition(ignorecondition bool) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) RecordExecution(recordexecution bool) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) SimulatedActions(simulatedactions types.SimulatedActionsVariant) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) TriggerData(triggerdata types.ScheduleTriggerEventVariant) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteWatch) Watch(watch types.WatchVariant) *ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}
