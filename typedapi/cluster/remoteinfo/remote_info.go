package remoteinfo

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RemoteInfo struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRemoteInfo func() *RemoteInfo

func NewRemoteInfoFunc(tp elastictransport.Interface) NewRemoteInfo {
	_ = "STUB: not implemented"
	return *new(NewRemoteInfo)
}

func New(tp elastictransport.Interface) *RemoteInfo { _ = "STUB: not implemented"; return nil }

func (r *RemoteInfo) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemoteInfo) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemoteInfo) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r RemoteInfo) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *RemoteInfo) Header(key, value string) *RemoteInfo { _ = "STUB: not implemented"; return nil }

func (r *RemoteInfo) ErrorTrace(errortrace bool) *RemoteInfo { _ = "STUB: not implemented"; return nil }

func (r *RemoteInfo) FilterPath(filterpaths ...string) *RemoteInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteInfo) Human(human bool) *RemoteInfo { _ = "STUB: not implemented"; return nil }

func (r *RemoteInfo) Pretty(pretty bool) *RemoteInfo { _ = "STUB: not implemented"; return nil }
