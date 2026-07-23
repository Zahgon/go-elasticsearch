package types

type MlInferenceTrainedModels struct {
	All_                          MlCounter                      `json:"_all"`
	Count                         *MlInferenceTrainedModelsCount `json:"count,omitempty"`
	EstimatedHeapMemoryUsageBytes *JobStatistics                 `json:"estimated_heap_memory_usage_bytes,omitempty"`
	EstimatedOperations           *JobStatistics                 `json:"estimated_operations,omitempty"`
	ModelSizeBytes                *JobStatistics                 `json:"model_size_bytes,omitempty"`
}

func NewMlInferenceTrainedModels() *MlInferenceTrainedModels { _ = "STUB: not implemented"; return nil }
