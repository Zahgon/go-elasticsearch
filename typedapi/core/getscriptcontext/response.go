package getscriptcontext

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Contexts []types.GetScriptContext `json:"contexts"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
