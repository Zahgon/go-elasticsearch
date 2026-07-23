package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _inferenceEndpoint struct {
	v *types.InferenceEndpoint
}

func NewInferenceEndpoint(service string) *_inferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceEndpoint) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *_inferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceEndpoint) Service(service string) *_inferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceEndpoint) ServiceSettings(servicesettings json.RawMessage) *_inferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceEndpoint) TaskSettings(tasksettings json.RawMessage) *_inferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceEndpoint) InferenceEndpointCaster() *types.InferenceEndpoint {
	_ = "STUB: not implemented"
	return nil
}
