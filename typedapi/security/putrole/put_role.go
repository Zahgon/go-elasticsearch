package putrole

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRole struct {
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

type NewPutRole func(name string) *PutRole

func NewPutRoleFunc(tp elastictransport.Interface) NewPutRole {
	_ = "STUB: not implemented"
	return *new(NewPutRole)
}

func New(tp elastictransport.Interface) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Raw(raw io.Reader) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Request(req *Request) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRole) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRole) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutRole) Header(key, value string) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) _name(name string) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Refresh(refresh refresh.Refresh) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) ErrorTrace(errortrace bool) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) FilterPath(filterpaths ...string) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Human(human bool) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Pretty(pretty bool) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Applications(applications ...types.ApplicationPrivilegesVariant) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) ApplicationsValues(applicationsvalues []types.ApplicationPrivileges) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) Description(description string) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) Global(global map[string]json.RawMessage) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) AddGlobal(key string, value json.RawMessage) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) Indices(indices ...types.IndicesPrivilegesVariant) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) IndicesValues(indicesvalues []types.IndicesPrivileges) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) Metadata(metadata types.MetadataVariant) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) RemoteCluster(remoteclusters ...types.RemoteClusterPrivilegesVariant) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) RemoteClusterValues(remoteclustervalues []types.RemoteClusterPrivileges) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) RemoteIndices(remoteindices ...types.RemoteIndicesPrivilegesVariant) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) RemoteIndicesValues(remoteindicesvalues []types.RemoteIndicesPrivileges) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) RunAs(runas ...string) *PutRole { _ = "STUB: not implemented"; return nil }

func (r *PutRole) TransientMetadata(transientmetadata map[string]json.RawMessage) *PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRole) AddTransientMetadatum(key string, value json.RawMessage) *PutRole {
	_ = "STUB: not implemented"
	return nil
}
