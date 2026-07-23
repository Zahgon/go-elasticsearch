package removeblock

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged bool                             `json:"acknowledged"`
	Indices      []types.RemoveIndicesBlockStatus `json:"indices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
