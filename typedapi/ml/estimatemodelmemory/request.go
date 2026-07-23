package estimatemodelmemory

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AnalysisConfig *types.AnalysisConfig `json:"analysis_config,omitempty"`

	MaxBucketCardinality map[string]int64 `json:"max_bucket_cardinality,omitempty"`

	OverallCardinality map[string]int64 `json:"overall_cardinality,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
