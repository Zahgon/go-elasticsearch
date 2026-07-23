package analyze

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Detail *types.AnalyzeDetail `json:"detail,omitempty"`
	Tokens []types.AnalyzeToken `json:"tokens,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
