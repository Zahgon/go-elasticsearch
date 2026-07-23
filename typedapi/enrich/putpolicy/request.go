package putpolicy

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	GeoMatch *types.EnrichPolicy `json:"geo_match,omitempty"`

	Match *types.EnrichPolicy `json:"match,omitempty"`

	Range *types.EnrichPolicy `json:"range,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
