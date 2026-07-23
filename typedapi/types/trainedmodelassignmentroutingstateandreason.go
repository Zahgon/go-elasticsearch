package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/routingstate"
)

type TrainedModelAssignmentRoutingStateAndReason struct {
	Reason *string `json:"reason,omitempty"`

	RoutingState routingstate.RoutingState `json:"routing_state"`
}

func (s *TrainedModelAssignmentRoutingStateAndReason) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelAssignmentRoutingStateAndReason() *TrainedModelAssignmentRoutingStateAndReason {
	_ = "STUB: not implemented"
	return nil
}
