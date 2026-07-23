package bulkupdateapikeys

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Expiration types.Duration `json:"expiration,omitempty"`

	Ids []string `json:"ids"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	RoleDescriptors map[string]types.RoleDescriptor `json:"role_descriptors,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
