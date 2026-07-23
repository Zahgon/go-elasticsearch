package segments

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Indices map[string]types.IndexSegment `json:"indices"`
	Shards_ types.ShardStatistics         `json:"_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
