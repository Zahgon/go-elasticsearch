package updatedatafeed

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Aggregations map[string]types.Aggregations `json:"aggregations,omitempty"`

	ChunkingConfig *types.ChunkingConfig `json:"chunking_config,omitempty"`

	DelayedDataCheckConfig *types.DelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`

	Frequency types.Duration `json:"frequency,omitempty"`

	Indices []string `json:"indices,omitempty"`

	IndicesOptions *types.IndicesOptions `json:"indices_options,omitempty"`
	JobId          *string               `json:"job_id,omitempty"`

	MaxEmptySearches *int `json:"max_empty_searches,omitempty"`

	Query *types.Query `json:"query,omitempty"`

	QueryDelay types.Duration `json:"query_delay,omitempty"`

	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`

	ScriptFields map[string]types.ScriptField `json:"script_fields,omitempty"`

	ScrollSize *int `json:"scroll_size,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
