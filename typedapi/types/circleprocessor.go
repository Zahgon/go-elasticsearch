package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shapetype"
)

type CircleProcessor struct {
	Description *string `json:"description,omitempty"`

	ErrorDistance Float64 `json:"error_distance"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	ShapeType shapetype.ShapeType `json:"shape_type"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *CircleProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCircleProcessor() *CircleProcessor { _ = "STUB: not implemented"; return nil }

type CircleProcessorVariant interface {
	CircleProcessorCaster() *CircleProcessor
}

func (s *CircleProcessor) CircleProcessorCaster() *CircleProcessor {
	_ = "STUB: not implemented"
	return nil
}
