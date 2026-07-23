package types

import (
	"encoding/json"
)

type InferenceEndpoint struct {
	ChunkingSettings *InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service string `json:"service"`

	ServiceSettings json.RawMessage `json:"service_settings"`

	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
}

func (s *InferenceEndpoint) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInferenceEndpoint() *InferenceEndpoint { _ = "STUB: not implemented"; return nil }

type InferenceEndpointVariant interface {
	InferenceEndpointCaster() *InferenceEndpoint
}

func (s *InferenceEndpoint) InferenceEndpointCaster() *InferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}
