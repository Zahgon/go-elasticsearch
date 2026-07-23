package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deploymentallocationstate"
)

type TrainedModelDeploymentAllocationStatus struct {
	AllocationCount int `json:"allocation_count"`

	State deploymentallocationstate.DeploymentAllocationState `json:"state"`

	TargetAllocationCount int `json:"target_allocation_count"`
}

func (s *TrainedModelDeploymentAllocationStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelDeploymentAllocationStatus() *TrainedModelDeploymentAllocationStatus {
	_ = "STUB: not implemented"
	return nil
}
