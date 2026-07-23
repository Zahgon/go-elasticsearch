package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deploymentassignmentstate"
)

type TrainedModelAssignment struct {
	AdaptiveAllocations *AdaptiveAllocationsSettings `json:"adaptive_allocations,omitempty"`

	AssignmentState        deploymentassignmentstate.DeploymentAssignmentState `json:"assignment_state"`
	MaxAssignedAllocations *int                                                `json:"max_assigned_allocations,omitempty"`
	Reason                 *string                                             `json:"reason,omitempty"`

	RoutingTable map[string]TrainedModelAssignmentRoutingTable `json:"routing_table"`

	StartTime      DateTime                             `json:"start_time"`
	TaskParameters TrainedModelAssignmentTaskParameters `json:"task_parameters"`
}

func (s *TrainedModelAssignment) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelAssignment() *TrainedModelAssignment { _ = "STUB: not implemented"; return nil }
