package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainingpriority"
)

type TrainedModelAssignmentTaskParameters struct {
	CacheSize ByteSize `json:"cache_size,omitempty"`

	DeploymentId string `json:"deployment_id"`

	ModelBytes ByteSize `json:"model_bytes"`

	ModelId string `json:"model_id"`

	NumberOfAllocations      int                               `json:"number_of_allocations"`
	PerAllocationMemoryBytes ByteSize                          `json:"per_allocation_memory_bytes"`
	PerDeploymentMemoryBytes ByteSize                          `json:"per_deployment_memory_bytes"`
	Priority                 trainingpriority.TrainingPriority `json:"priority"`

	QueueCapacity int `json:"queue_capacity"`

	ThreadsPerAllocation int `json:"threads_per_allocation"`
}

func (s *TrainedModelAssignmentTaskParameters) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelAssignmentTaskParameters() *TrainedModelAssignmentTaskParameters {
	_ = "STUB: not implemented"
	return nil
}
