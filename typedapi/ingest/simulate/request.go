package simulate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Docs []types.Document `json:"docs"`

	Pipeline *types.IngestPipeline `json:"pipeline,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
