package types

import (
	"encoding/json"
)

type CsvProcessor struct {
	Description *string `json:"description,omitempty"`

	EmptyValue json.RawMessage `json:"empty_value,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Quote *string `json:"quote,omitempty"`

	Separator *string `json:"separator,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetFields []string `json:"target_fields"`

	Trim *bool `json:"trim,omitempty"`
}

func (s *CsvProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCsvProcessor() *CsvProcessor { _ = "STUB: not implemented"; return nil }

type CsvProcessorVariant interface {
	CsvProcessorCaster() *CsvProcessor
}

func (s *CsvProcessor) CsvProcessorCaster() *CsvProcessor { _ = "STUB: not implemented"; return nil }
