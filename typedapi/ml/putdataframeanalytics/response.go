package putdataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AllowLazyStart   bool                                   `json:"allow_lazy_start"`
	Analysis         types.DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *types.DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	Authorization    *types.DataframeAnalyticsAuthorization `json:"authorization,omitempty"`
	CreateTime       int64                                  `json:"create_time"`
	Description      *string                                `json:"description,omitempty"`
	Dest             types.DataframeAnalyticsDestination    `json:"dest"`
	Id               string                                 `json:"id"`
	MaxNumThreads    int                                    `json:"max_num_threads"`
	Meta_            types.Metadata                         `json:"_meta,omitempty"`
	ModelMemoryLimit string                                 `json:"model_memory_limit"`
	Source           types.DataframeAnalyticsSource         `json:"source"`
	Version          string                                 `json:"version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
