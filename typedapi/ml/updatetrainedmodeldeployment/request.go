package updatetrainedmodeldeployment

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AdaptiveAllocations *types.AdaptiveAllocationsSettings `json:"adaptive_allocations,omitempty"`

	NumberOfAllocations *int `json:"number_of_allocations,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
