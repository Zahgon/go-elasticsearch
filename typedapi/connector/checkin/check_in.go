package checkin

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CheckIn struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCheckIn func(connectorid string) *CheckIn

func NewCheckInFunc(tp elastictransport.Interface) NewCheckIn {
	_ = "STUB: not implemented"
	return *new(NewCheckIn)
}

func New(tp elastictransport.Interface) *CheckIn { _ = "STUB: not implemented"; return nil }

func (r *CheckIn) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CheckIn) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CheckIn) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CheckIn) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CheckIn) Header(key, value string) *CheckIn { _ = "STUB: not implemented"; return nil }

func (r *CheckIn) _connectorid(connectorid string) *CheckIn { _ = "STUB: not implemented"; return nil }

func (r *CheckIn) ErrorTrace(errortrace bool) *CheckIn { _ = "STUB: not implemented"; return nil }

func (r *CheckIn) FilterPath(filterpaths ...string) *CheckIn { _ = "STUB: not implemented"; return nil }

func (r *CheckIn) Human(human bool) *CheckIn { _ = "STUB: not implemented"; return nil }

func (r *CheckIn) Pretty(pretty bool) *CheckIn { _ = "STUB: not implemented"; return nil }
