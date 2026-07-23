package types

import (
	"encoding/json"
)

type ModelPackageConfig struct {
	CreateTime           *int64                     `json:"create_time,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	InferenceConfig      map[string]json.RawMessage `json:"inference_config,omitempty"`
	Metadata             Metadata                   `json:"metadata,omitempty"`
	MinimumVersion       *string                    `json:"minimum_version,omitempty"`
	ModelRepository      *string                    `json:"model_repository,omitempty"`
	ModelType            *string                    `json:"model_type,omitempty"`
	PackagedModelId      string                     `json:"packaged_model_id"`
	PlatformArchitecture *string                    `json:"platform_architecture,omitempty"`
	PrefixStrings        *TrainedModelPrefixStrings `json:"prefix_strings,omitempty"`
	Sha256               *string                    `json:"sha256,omitempty"`
	Size                 ByteSize                   `json:"size,omitempty"`
	Tags                 []string                   `json:"tags,omitempty"`
	VocabularyFile       *string                    `json:"vocabulary_file,omitempty"`
}

func (s *ModelPackageConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewModelPackageConfig() *ModelPackageConfig { _ = "STUB: not implemented"; return nil }
