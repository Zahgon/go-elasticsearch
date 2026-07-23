package rollupsearch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Aggregations    map[string]types.Aggregate `json:"aggregations,omitempty"`
	Hits            types.HitsMetadata         `json:"hits"`
	Shards_         types.ShardStatistics      `json:"_shards"`
	TerminatedEarly *bool                      `json:"terminated_early,omitempty"`
	TimedOut        bool                       `json:"timed_out"`
	Took            int64                      `json:"took"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
