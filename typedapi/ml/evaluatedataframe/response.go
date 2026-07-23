package evaluatedataframe

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Classification *types.DataframeClassificationSummary `json:"classification,omitempty"`

	OutlierDetection *types.DataframeOutlierDetectionSummary `json:"outlier_detection,omitempty"`

	Regression *types.DataframeRegressionSummary `json:"regression,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
