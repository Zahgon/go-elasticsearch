package explaindataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	FieldSelection []types.DataframeAnalyticsFieldSelection `json:"field_selection"`

	MemoryEstimation types.DataframeAnalyticsMemoryEstimation `json:"memory_estimation"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
