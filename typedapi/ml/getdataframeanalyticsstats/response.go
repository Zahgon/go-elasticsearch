package getdataframeanalyticsstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int64 `json:"count"`

	DataFrameAnalytics []types.DataframeAnalytics `json:"data_frame_analytics"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
