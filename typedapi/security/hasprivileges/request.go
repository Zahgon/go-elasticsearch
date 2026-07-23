package hasprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type Request struct {
	Application []types.ApplicationPrivilegesCheck `json:"application,omitempty"`

	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	Index   []types.IndexPrivilegesCheck        `json:"index,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
