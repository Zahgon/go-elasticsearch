package explaindataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowLazyStart *bool `json:"allow_lazy_start,omitempty"`

	Analysis *types.DataframeAnalysisContainer `json:"analysis,omitempty"`

	AnalyzedFields *types.DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`

	Description *string `json:"description,omitempty"`

	Dest *types.DataframeAnalyticsDestination `json:"dest,omitempty"`

	MaxNumThreads *int `json:"max_num_threads,omitempty"`

	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`

	Source *types.DataframeAnalyticsSource `json:"source,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
