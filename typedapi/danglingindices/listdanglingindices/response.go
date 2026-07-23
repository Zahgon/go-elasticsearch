package listdanglingindices

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	DanglingIndices []types.DanglingIndex `json:"dangling_indices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
