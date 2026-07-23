package putnode

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/type_"
)

type Request struct {
	AllocationDelay *string `json:"allocation_delay,omitempty"`

	Reason string `json:"reason"`

	TargetNodeName *string `json:"target_node_name,omitempty"`

	Type type_.Type `json:"type"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
