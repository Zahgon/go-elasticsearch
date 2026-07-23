package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type FloatRangeProperty struct {
	Boost       *Float64                       `json:"boost,omitempty"`
	Coerce      *bool                          `json:"coerce,omitempty"`
	CopyTo      []string                       `json:"copy_to,omitempty"`
	DocValues   *bool                          `json:"doc_values,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`
	Index       *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *FloatRangeProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FloatRangeProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFloatRangeProperty() *FloatRangeProperty { _ = "STUB: not implemented"; return nil }

type FloatRangePropertyVariant interface {
	FloatRangePropertyCaster() *FloatRangeProperty
}

func (s *FloatRangeProperty) FloatRangePropertyCaster() *FloatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *FloatRangeProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
