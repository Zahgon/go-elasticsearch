package types

import (
	"encoding/json"
)

type TrainedModelStats struct {
	DeploymentStats *TrainedModelDeploymentStats `json:"deployment_stats,omitempty"`

	InferenceStats *TrainedModelInferenceStats `json:"inference_stats,omitempty"`

	Ingest map[string]json.RawMessage `json:"ingest,omitempty"`

	ModelId string `json:"model_id"`

	ModelSizeStats TrainedModelSizeStats `json:"model_size_stats"`

	PipelineCount int `json:"pipeline_count"`
}

func (s *TrainedModelStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTrainedModelStats() *TrainedModelStats { _ = "STUB: not implemented"; return nil }
