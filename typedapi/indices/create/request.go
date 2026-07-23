package create

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Aliases map[string]types.Alias `json:"aliases,omitempty"`

	Mappings *types.TypeMapping `json:"mappings,omitempty"`

	Settings *types.IndexSettings `json:"settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
