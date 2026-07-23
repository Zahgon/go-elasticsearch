package restore

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	FeatureStates []string `json:"feature_states,omitempty"`

	IgnoreIndexSettings []string `json:"ignore_index_settings,omitempty"`

	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`

	IncludeAliases *bool `json:"include_aliases,omitempty"`

	IncludeGlobalState *bool `json:"include_global_state,omitempty"`

	IndexSettings *types.IndexSettings `json:"index_settings,omitempty"`

	Indices []string `json:"indices,omitempty"`

	Partial *bool `json:"partial,omitempty"`

	RenamePattern *string `json:"rename_pattern,omitempty"`

	RenameReplacement *string `json:"rename_replacement,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
