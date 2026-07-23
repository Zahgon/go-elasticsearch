package hasprivilegesuserprofile

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HasPrivilegesUserProfile struct {
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

type NewHasPrivilegesUserProfile func() *HasPrivilegesUserProfile

func NewHasPrivilegesUserProfileFunc(tp elastictransport.Interface) NewHasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return *new(NewHasPrivilegesUserProfile)
}

func New(tp elastictransport.Interface) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) Raw(raw io.Reader) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) Request(req *Request) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HasPrivilegesUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HasPrivilegesUserProfile) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *HasPrivilegesUserProfile) Header(key, value string) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) ErrorTrace(errortrace bool) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) FilterPath(filterpaths ...string) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) Human(human bool) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) Pretty(pretty bool) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) Privileges(privileges types.PrivilegesCheckVariant) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivilegesUserProfile) Uids(uids ...string) *HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}
