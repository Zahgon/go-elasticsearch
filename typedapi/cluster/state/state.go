package state

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
	metricMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type State struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	metric string
	index  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewState func() *State

func NewStateFunc(tp elastictransport.Interface) NewState {
	_ = "STUB: not implemented"
	return *new(NewState)
}

func New(tp elastictransport.Interface) *State { _ = "STUB: not implemented"; return nil }

func (r *State) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r State) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r State) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r State) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *State) Header(key, value string) *State { _ = "STUB: not implemented"; return nil }

func (r *State) Metric(metric string) *State { _ = "STUB: not implemented"; return nil }

func (r *State) Index(index string) *State { _ = "STUB: not implemented"; return nil }

func (r *State) AllowNoIndices(allownoindices bool) *State { _ = "STUB: not implemented"; return nil }

func (r *State) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *State {
	_ = "STUB: not implemented"
	return nil
}

func (r *State) FlatSettings(flatsettings bool) *State { _ = "STUB: not implemented"; return nil }

func (r *State) IgnoreUnavailable(ignoreunavailable bool) *State {
	_ = "STUB: not implemented"
	return nil
}

func (r *State) Local(local bool) *State { _ = "STUB: not implemented"; return nil }

func (r *State) MasterTimeout(duration string) *State { _ = "STUB: not implemented"; return nil }

func (r *State) WaitForMetadataVersion(versionnumber string) *State {
	_ = "STUB: not implemented"
	return nil
}

func (r *State) WaitForTimeout(duration string) *State { _ = "STUB: not implemented"; return nil }

func (r *State) ErrorTrace(errortrace bool) *State { _ = "STUB: not implemented"; return nil }

func (r *State) FilterPath(filterpaths ...string) *State { _ = "STUB: not implemented"; return nil }

func (r *State) Human(human bool) *State { _ = "STUB: not implemented"; return nil }

func (r *State) Pretty(pretty bool) *State { _ = "STUB: not implemented"; return nil }
