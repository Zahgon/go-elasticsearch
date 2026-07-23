package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type CountedKeywordProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`
	Index       *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *CountedKeywordProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CountedKeywordProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCountedKeywordProperty() *CountedKeywordProperty { _ = "STUB: not implemented"; return nil }

type CountedKeywordPropertyVariant interface {
	CountedKeywordPropertyCaster() *CountedKeywordProperty
}

func (s *CountedKeywordProperty) CountedKeywordPropertyCaster() *CountedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *CountedKeywordProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
