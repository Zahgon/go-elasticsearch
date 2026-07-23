package getdataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	DataFrameAnalytics []types.DataframeAnalyticsSummary `json:"data_frame_analytics"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
