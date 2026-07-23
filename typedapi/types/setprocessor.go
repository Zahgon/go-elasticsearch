package types

import (
	"encoding/json"
)

type SetProcessor struct {
	CopyFrom *string `json:"copy_from,omitempty"`

	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreEmptyValue *bool `json:"ignore_empty_value,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	MediaType *string `json:"media_type,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Override *bool `json:"override,omitempty"`

	Tag *string `json:"tag,omitempty"`

	Value json.RawMessage `json:"value,omitempty"`
}

func (s *SetProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSetProcessor() *SetProcessor { _ = "STUB: not implemented"; return nil }

type SetProcessorVariant interface {
	SetProcessorCaster() *SetProcessor
}

func (s *SetProcessor) SetProcessorCaster() *SetProcessor { _ = "STUB: not implemented"; return nil }
