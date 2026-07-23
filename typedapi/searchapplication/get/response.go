package get

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AnalyticsCollectionName *string `json:"analytics_collection_name,omitempty"`

	Indices []string `json:"indices"`

	Name string `json:"name"`

	Template *types.SearchApplicationTemplate `json:"template,omitempty"`

	UpdatedAtMillis int64 `json:"updated_at_millis"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
