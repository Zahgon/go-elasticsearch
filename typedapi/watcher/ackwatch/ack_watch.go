package ackwatch

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	watchidMask = iota + 1

	actionidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AckWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	watchid  string
	actionid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewAckWatch func(watchid string) *AckWatch

func NewAckWatchFunc(tp elastictransport.Interface) NewAckWatch {
	_ = "STUB: not implemented"
	return *new(NewAckWatch)
}

func New(tp elastictransport.Interface) *AckWatch { _ = "STUB: not implemented"; return nil }

func (r *AckWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AckWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AckWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AckWatch) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *AckWatch) Header(key, value string) *AckWatch { _ = "STUB: not implemented"; return nil }

func (r *AckWatch) _watchid(watchid string) *AckWatch { _ = "STUB: not implemented"; return nil }

func (r *AckWatch) ActionId(actionid string) *AckWatch { _ = "STUB: not implemented"; return nil }

func (r *AckWatch) ErrorTrace(errortrace bool) *AckWatch { _ = "STUB: not implemented"; return nil }

func (r *AckWatch) FilterPath(filterpaths ...string) *AckWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *AckWatch) Human(human bool) *AckWatch { _ = "STUB: not implemented"; return nil }

func (r *AckWatch) Pretty(pretty bool) *AckWatch { _ = "STUB: not implemented"; return nil }
