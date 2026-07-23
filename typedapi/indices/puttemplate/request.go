package puttemplate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Aliases map[string]types.Alias `json:"aliases,omitempty"`

	IndexPatterns []string `json:"index_patterns,omitempty"`

	Mappings *types.TypeMapping `json:"mappings,omitempty"`

	Order *int `json:"order,omitempty"`

	Settings *types.IndexSettings `json:"settings,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
