package hasprivileges

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

const (
	userMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HasPrivileges struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	user string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewHasPrivileges func() *HasPrivileges

func NewHasPrivilegesFunc(tp elastictransport.Interface) NewHasPrivileges {
	_ = "STUB: not implemented"
	return *new(NewHasPrivileges)
}

func New(tp elastictransport.Interface) *HasPrivileges { _ = "STUB: not implemented"; return nil }

func (r *HasPrivileges) Raw(raw io.Reader) *HasPrivileges { _ = "STUB: not implemented"; return nil }

func (r *HasPrivileges) Request(req *Request) *HasPrivileges { _ = "STUB: not implemented"; return nil }

func (r *HasPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HasPrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HasPrivileges) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *HasPrivileges) Header(key, value string) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) User(user string) *HasPrivileges { _ = "STUB: not implemented"; return nil }

func (r *HasPrivileges) ErrorTrace(errortrace bool) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) FilterPath(filterpaths ...string) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) Human(human bool) *HasPrivileges { _ = "STUB: not implemented"; return nil }

func (r *HasPrivileges) Pretty(pretty bool) *HasPrivileges { _ = "STUB: not implemented"; return nil }

func (r *HasPrivileges) Application(applications ...types.ApplicationPrivilegesCheckVariant) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) ApplicationValues(applicationvalues []types.ApplicationPrivilegesCheck) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) Index(indices ...types.IndexPrivilegesCheckVariant) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *HasPrivileges) IndexValues(indexvalues []types.IndexPrivilegesCheck) *HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}
