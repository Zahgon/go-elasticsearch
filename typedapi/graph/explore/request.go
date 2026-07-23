package explore

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Connections *types.Hop `json:"connections,omitempty"`

	Controls *types.ExploreControls `json:"controls,omitempty"`

	Query *types.Query `json:"query,omitempty"`

	Vertices []types.VertexDefinition `json:"vertices,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
