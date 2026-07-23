package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tasktype"
)

type InferenceEndpointInfo struct {
	ChunkingSettings *InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	InferenceId string `json:"inference_id"`

	Service string `json:"service"`

	ServiceSettings json.RawMessage `json:"service_settings"`

	TaskSettings json.RawMessage `json:"task_settings,omitempty"`

	TaskType tasktype.TaskType `json:"task_type"`
}

func (s *InferenceEndpointInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceEndpointInfo() *InferenceEndpointInfo { _ = "STUB: not implemented"; return nil }
