package post

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	License *types.License `json:"license,omitempty"`

	Licenses []types.License `json:"licenses,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
