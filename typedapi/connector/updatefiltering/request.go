package updatefiltering

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AdvancedSnippet *types.FilteringAdvancedSnippet `json:"advanced_snippet,omitempty"`
	Filtering       []types.FilteringConfig         `json:"filtering,omitempty"`
	Rules           []types.FilteringRule           `json:"rules,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
