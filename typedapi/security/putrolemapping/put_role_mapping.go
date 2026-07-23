package putrolemapping

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRoleMapping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutRoleMapping func(name string) *PutRoleMapping

func NewPutRoleMappingFunc(tp elastictransport.Interface) NewPutRoleMapping {
	_ = "STUB: not implemented"
	return *new(NewPutRoleMapping)
}

func New(tp elastictransport.Interface) *PutRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *PutRoleMapping) Raw(raw io.Reader) *PutRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *PutRoleMapping) Request(req *Request) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRoleMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRoleMapping) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutRoleMapping) Header(key, value string) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) _name(name string) *PutRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *PutRoleMapping) Refresh(refresh refresh.Refresh) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) ErrorTrace(errortrace bool) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) FilterPath(filterpaths ...string) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) Human(human bool) *PutRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *PutRoleMapping) Pretty(pretty bool) *PutRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *PutRoleMapping) Enabled(enabled bool) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) Metadata(metadata types.MetadataVariant) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) RoleTemplates(roletemplates ...types.RoleTemplateVariant) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) RoleTemplatesValues(roletemplatesvalues []types.RoleTemplate) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) Roles(roles ...string) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) Rules(rules types.RoleMappingRuleVariant) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRoleMapping) RunAs(runas ...string) *PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}
