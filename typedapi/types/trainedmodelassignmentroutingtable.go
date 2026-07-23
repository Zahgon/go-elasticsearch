package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/routingstate"
)

type TrainedModelAssignmentRoutingTable struct {
	CurrentAllocations int `json:"current_allocations"`

	Reason *string `json:"reason,omitempty"`

	RoutingState routingstate.RoutingState `json:"routing_state"`

	TargetAllocations int `json:"target_allocations"`
}

func (s *TrainedModelAssignmentRoutingTable) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelAssignmentRoutingTable() *TrainedModelAssignmentRoutingTable {
	_ = "STUB: not implemented"
	return nil
}
