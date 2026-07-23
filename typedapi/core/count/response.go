package count

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count   int64                 `json:"count"`
	Shards_ types.ShardStatistics `json:"_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
