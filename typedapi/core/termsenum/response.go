package termsenum

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Complete bool                  `json:"complete"`
	Shards_  types.ShardStatistics `json:"_shards"`
	Terms    []string              `json:"terms"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
