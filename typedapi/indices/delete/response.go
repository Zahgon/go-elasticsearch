package delete

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged bool                   `json:"acknowledged"`
	Shards_      *types.ShardStatistics `json:"_shards,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
