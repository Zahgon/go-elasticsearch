package verifyrepository

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Nodes map[string]types.CompactNodeInfo `json:"nodes"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
