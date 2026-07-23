package knnsearch

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Fields map[string]json.RawMessage `json:"fields,omitempty"`

	Hits types.HitsMetadata `json:"hits"`

	MaxScore *types.Float64 `json:"max_score,omitempty"`

	Shards_ types.ShardStatistics `json:"_shards"`

	TimedOut bool `json:"timed_out"`

	Took int64 `json:"took"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
