package reloadsearchanalyzers

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ReloadDetails []types.ReloadDetails `json:"reload_details"`
	Shards_       types.ShardStatistics `json:"_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
