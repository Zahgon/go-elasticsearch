package hasprivilegesuserprofile

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Errors *types.HasPrivilegesUserProfileErrors `json:"errors,omitempty"`

	HasPrivilegeUids []string `json:"has_privilege_uids"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
