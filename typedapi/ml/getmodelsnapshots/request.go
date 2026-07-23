package getmodelsnapshots

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Desc *bool `json:"desc,omitempty"`

	End  types.DateTime `json:"end,omitempty"`
	Page *types.Page    `json:"page,omitempty"`

	Sort *string `json:"sort,omitempty"`

	Start types.DateTime `json:"start,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
