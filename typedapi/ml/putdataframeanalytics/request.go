package putdataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowLazyStart *bool `json:"allow_lazy_start,omitempty"`

	Analysis types.DataframeAnalysisContainer `json:"analysis"`

	AnalyzedFields *types.DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`

	Description *string `json:"description,omitempty"`

	Dest    types.DataframeAnalyticsDestination `json:"dest"`
	Headers types.HttpHeaders                   `json:"headers,omitempty"`

	MaxNumThreads *int           `json:"max_num_threads,omitempty"`
	Meta_         types.Metadata `json:"_meta,omitempty"`

	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`

	Source  types.DataframeAnalyticsSource `json:"source"`
	Version *string                        `json:"version,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
