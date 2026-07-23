package info

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Defaults    types.Defaults   `json:"defaults"`
	Limits      types.Limits     `json:"limits"`
	NativeCode  types.NativeCode `json:"native_code"`
	UpgradeMode bool             `json:"upgrade_mode"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
