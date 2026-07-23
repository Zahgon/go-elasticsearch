package putprivileges

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutPrivileges struct {
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

type NewPutPrivileges func() *PutPrivileges

func NewPutPrivilegesFunc(tp elastictransport.Interface) NewPutPrivileges {
	_ = "STUB: not implemented"
	return *new(NewPutPrivileges)
}

func New(tp elastictransport.Interface) *PutPrivileges { _ = "STUB: not implemented"; return nil }

func (r *PutPrivileges) Raw(raw io.Reader) *PutPrivileges { _ = "STUB: not implemented"; return nil }

func (r *PutPrivileges) Request(req *Request) *PutPrivileges { _ = "STUB: not implemented"; return nil }

func (r *PutPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutPrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutPrivileges) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *PutPrivileges) Header(key, value string) *PutPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPrivileges) Refresh(refresh refresh.Refresh) *PutPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPrivileges) ErrorTrace(errortrace bool) *PutPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPrivileges) FilterPath(filterpaths ...string) *PutPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPrivileges) Human(human bool) *PutPrivileges { _ = "STUB: not implemented"; return nil }

func (r *PutPrivileges) Pretty(pretty bool) *PutPrivileges { _ = "STUB: not implemented"; return nil }
