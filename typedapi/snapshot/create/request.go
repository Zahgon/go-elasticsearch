package create

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

type Request struct {
	ExpandWildcards []expandwildcard.ExpandWildcard `json:"expand_wildcards,omitempty"`

	FeatureStates []string `json:"feature_states,omitempty"`

	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`

	IncludeGlobalState *bool `json:"include_global_state,omitempty"`

	Indices []string `json:"indices,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	Partial *bool `json:"partial,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
