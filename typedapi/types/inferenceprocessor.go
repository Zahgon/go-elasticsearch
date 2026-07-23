package types

import (
	"encoding/json"
)

type InferenceProcessor struct {
	Description *string `json:"description,omitempty"`

	FieldMap map[string]json.RawMessage `json:"field_map,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	InferenceConfig *InferenceConfig `json:"inference_config,omitempty"`

	InputOutput []InputConfig `json:"input_output,omitempty"`

	ModelId string `json:"model_id"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *InferenceProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceProcessor() *InferenceProcessor { _ = "STUB: not implemented"; return nil }

type InferenceProcessorVariant interface {
	InferenceProcessorCaster() *InferenceProcessor
}

func (s *InferenceProcessor) InferenceProcessorCaster() *InferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}
