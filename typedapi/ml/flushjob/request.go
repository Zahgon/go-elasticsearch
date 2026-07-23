package flushjob

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AdvanceTime types.DateTime `json:"advance_time,omitempty"`

	CalcInterim *bool `json:"calc_interim,omitempty"`

	End types.DateTime `json:"end,omitempty"`

	SkipTime types.DateTime `json:"skip_time,omitempty"`

	Start types.DateTime `json:"start,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
