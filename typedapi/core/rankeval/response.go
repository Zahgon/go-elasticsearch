package rankeval

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Details  map[string]types.RankEvalMetricDetail `json:"details"`
	Failures map[string]json.RawMessage            `json:"failures"`

	MetricScore types.Float64 `json:"metric_score"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
