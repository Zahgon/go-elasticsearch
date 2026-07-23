package healthreport

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indicatorhealthstatus"
)

type Response struct {
	ClusterName string                                       `json:"cluster_name"`
	Indicators  types.Indicators                             `json:"indicators"`
	Status      *indicatorhealthstatus.IndicatorHealthStatus `json:"status,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
