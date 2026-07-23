package query

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowPartialSearchResults *bool `json:"allow_partial_search_results,omitempty"`

	Catalog *string `json:"catalog,omitempty"`

	Columnar *bool `json:"columnar,omitempty"`

	Cursor *string `json:"cursor,omitempty"`

	FetchSize *int `json:"fetch_size,omitempty"`

	FieldMultiValueLeniency *bool `json:"field_multi_value_leniency,omitempty"`

	Filter *types.Query `json:"filter,omitempty"`

	IndexUsingFrozen *bool `json:"index_using_frozen,omitempty"`

	KeepAlive types.Duration `json:"keep_alive,omitempty"`

	KeepOnCompletion *bool `json:"keep_on_completion,omitempty"`

	PageTimeout types.Duration `json:"page_timeout,omitempty"`

	Params []json.RawMessage `json:"params,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query *string `json:"query,omitempty"`

	RequestTimeout types.Duration `json:"request_timeout,omitempty"`

	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`

	WaitForCompletionTimeout types.Duration `json:"wait_for_completion_timeout,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
