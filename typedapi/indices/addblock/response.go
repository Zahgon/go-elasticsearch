package addblock

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged       bool                          `json:"acknowledged"`
	Indices            []types.AddIndicesBlockStatus `json:"indices"`
	ShardsAcknowledged bool                          `json:"shards_acknowledged"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
