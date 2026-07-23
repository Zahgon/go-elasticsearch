package search

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/resultposition"
)

type Request struct {
	AllowPartialSearchResults *bool `json:"allow_partial_search_results,omitempty"`

	AllowPartialSequenceResults *bool `json:"allow_partial_sequence_results,omitempty"`
	CaseSensitive               *bool `json:"case_sensitive,omitempty"`

	EventCategoryField *string `json:"event_category_field,omitempty"`

	FetchSize *uint `json:"fetch_size,omitempty"`

	Fields []types.FieldAndFormat `json:"fields,omitempty"`

	Filter           []types.Query  `json:"filter,omitempty"`
	KeepAlive        types.Duration `json:"keep_alive,omitempty"`
	KeepOnCompletion *bool          `json:"keep_on_completion,omitempty"`

	MaxSamplesPerKey *int `json:"max_samples_per_key,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query           string                         `json:"query"`
	ResultPosition  *resultposition.ResultPosition `json:"result_position,omitempty"`
	RuntimeMappings types.RuntimeFields            `json:"runtime_mappings,omitempty"`

	Size *uint `json:"size,omitempty"`

	TiebreakerField *string `json:"tiebreaker_field,omitempty"`

	TimestampField           *string        `json:"timestamp_field,omitempty"`
	WaitForCompletionTimeout types.Duration `json:"wait_for_completion_timeout,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
