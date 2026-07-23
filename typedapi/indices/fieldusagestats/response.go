package fieldusagestats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	FieldsUsageBody map[string]types.UsageStatsIndex `json:"-"`
	Shards_         types.ShardStatistics            `json:"_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
