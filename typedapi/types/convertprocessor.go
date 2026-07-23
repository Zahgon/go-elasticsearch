package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/converttype"
)

type ConvertProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	Type converttype.ConvertType `json:"type"`
}

func (s *ConvertProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewConvertProcessor() *ConvertProcessor { _ = "STUB: not implemented"; return nil }

type ConvertProcessorVariant interface {
	ConvertProcessorCaster() *ConvertProcessor
}

func (s *ConvertProcessor) ConvertProcessorCaster() *ConvertProcessor {
	_ = "STUB: not implemented"
	return nil
}
