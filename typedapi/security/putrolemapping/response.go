package putrolemapping

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Created     *bool               `json:"created,omitempty"`
	RoleMapping types.CreatedStatus `json:"role_mapping"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
