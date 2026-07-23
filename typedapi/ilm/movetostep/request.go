package movetostep

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	CurrentStep types.StepKey `json:"current_step"`

	NextStep types.StepKey `json:"next_step"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
