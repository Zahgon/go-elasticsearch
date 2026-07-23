package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deploymentassignmentstate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainingpriority"
)

type TrainedModelDeploymentStats struct {
	AdaptiveAllocations *AdaptiveAllocationsSettings `json:"adaptive_allocations,omitempty"`

	AllocationStatus *TrainedModelDeploymentAllocationStatus `json:"allocation_status,omitempty"`
	CacheSize        ByteSize                                `json:"cache_size,omitempty"`

	DeploymentId string `json:"deployment_id"`

	ErrorCount *int `json:"error_count,omitempty"`

	InferenceCount *int `json:"inference_count,omitempty"`

	ModelId string `json:"model_id"`

	Nodes []TrainedModelDeploymentNodesStats `json:"nodes"`

	NumberOfAllocations     *int                              `json:"number_of_allocations,omitempty"`
	PeakThroughputPerMinute int64                             `json:"peak_throughput_per_minute"`
	Priority                trainingpriority.TrainingPriority `json:"priority"`

	QueueCapacity *int `json:"queue_capacity,omitempty"`

	Reason *string `json:"reason,omitempty"`

	RejectedExecutionCount *int `json:"rejected_execution_count,omitempty"`

	StartTime int64 `json:"start_time"`

	State *deploymentassignmentstate.DeploymentAssignmentState `json:"state,omitempty"`

	ThreadsPerAllocation *int `json:"threads_per_allocation,omitempty"`

	TimeoutCount *int `json:"timeout_count,omitempty"`
}

func (s *TrainedModelDeploymentStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelDeploymentStats() *TrainedModelDeploymentStats {
	_ = "STUB: not implemented"
	return nil
}
