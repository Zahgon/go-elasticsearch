package rankeval

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Metric *types.RankEvalMetric `json:"metric,omitempty"`

	Requests []types.RankEvalRequestItem `json:"requests"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
