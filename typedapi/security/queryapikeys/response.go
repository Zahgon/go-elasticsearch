package queryapikeys

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Aggregations map[string]types.ApiKeyAggregate `json:"aggregations,omitempty"`

	ApiKeys []types.ApiKey `json:"api_keys"`

	Count int `json:"count"`

	Total int `json:"total"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
