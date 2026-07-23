package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type CompletionProperty struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Contexts       []SuggestContext               `json:"contexts,omitempty"`
	CopyTo         []string                       `json:"copy_to,omitempty"`
	DocValues      *bool                          `json:"doc_values,omitempty"`
	Dynamic        *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields         map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove    *int                           `json:"ignore_above,omitempty"`
	MaxInputLength *int                           `json:"max_input_length,omitempty"`

	Meta                       map[string]string                                `json:"meta,omitempty"`
	PreservePositionIncrements *bool                                            `json:"preserve_position_increments,omitempty"`
	PreserveSeparators         *bool                                            `json:"preserve_separators,omitempty"`
	Properties                 map[string]Property                              `json:"properties,omitempty"`
	SearchAnalyzer             *string                                          `json:"search_analyzer,omitempty"`
	Store                      *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep        *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                       string                                           `json:"type,omitempty"`
}

func (s *CompletionProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CompletionProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCompletionProperty() *CompletionProperty { _ = "STUB: not implemented"; return nil }

type CompletionPropertyVariant interface {
	CompletionPropertyCaster() *CompletionProperty
}

func (s *CompletionProperty) CompletionPropertyCaster() *CompletionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompletionProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
