package getsettings

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Security types.SecuritySettings `json:"security"`

	SecurityProfile types.SecuritySettings `json:"security-profile"`

	SecurityTokens types.SecuritySettings `json:"security-tokens"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
