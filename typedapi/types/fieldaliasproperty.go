package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type FieldAliasProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Path                *string                                          `json:"path,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *FieldAliasProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FieldAliasProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFieldAliasProperty() *FieldAliasProperty { _ = "STUB: not implemented"; return nil }

type FieldAliasPropertyVariant interface {
	FieldAliasPropertyCaster() *FieldAliasProperty
}

func (s *FieldAliasProperty) FieldAliasPropertyCaster() *FieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *FieldAliasProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
