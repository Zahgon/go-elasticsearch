package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type ConstantKeywordProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
	Value               json.RawMessage                                  `json:"value,omitempty"`
}

func (s *ConstantKeywordProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ConstantKeywordProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewConstantKeywordProperty() *ConstantKeywordProperty { _ = "STUB: not implemented"; return nil }

type ConstantKeywordPropertyVariant interface {
	ConstantKeywordPropertyCaster() *ConstantKeywordProperty
}

func (s *ConstantKeywordProperty) ConstantKeywordPropertyCaster() *ConstantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *ConstantKeywordProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
