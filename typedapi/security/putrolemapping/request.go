package putrolemapping

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Enabled *bool `json:"enabled,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	RoleTemplates []types.RoleTemplate `json:"role_templates,omitempty"`

	Roles []string `json:"roles,omitempty"`

	Rules *types.RoleMappingRule `json:"rules,omitempty"`
	RunAs []string               `json:"run_as,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
