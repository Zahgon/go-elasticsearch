package getrecords

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count   int64           `json:"count"`
	Records []types.Anomaly `json:"records"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
