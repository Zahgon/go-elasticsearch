package types

import (
	"encoding/json"
)

type AppendProcessor struct {
	AllowDuplicates *bool `json:"allow_duplicates,omitempty"`

	CopyFrom *string `json:"copy_from,omitempty"`

	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreEmptyValues *bool `json:"ignore_empty_values,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	MediaType *string `json:"media_type,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	Value []json.RawMessage `json:"value,omitempty"`
}

func (s *AppendProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAppendProcessor() *AppendProcessor { _ = "STUB: not implemented"; return nil }

type AppendProcessorVariant interface {
	AppendProcessorCaster() *AppendProcessor
}

func (s *AppendProcessor) AppendProcessorCaster() *AppendProcessor {
	_ = "STUB: not implemented"
	return nil
}
