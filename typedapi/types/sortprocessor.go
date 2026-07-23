package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type SortProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Order *sortorder.SortOrder `json:"order,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *SortProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSortProcessor() *SortProcessor { _ = "STUB: not implemented"; return nil }

type SortProcessorVariant interface {
	SortProcessorCaster() *SortProcessor
}

func (s *SortProcessor) SortProcessorCaster() *SortProcessor { _ = "STUB: not implemented"; return nil }
