package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type NestedProperty struct {
	CopyTo          []string                       `json:"copy_to,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled         *bool                          `json:"enabled,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IncludeInParent *bool                          `json:"include_in_parent,omitempty"`
	IncludeInRoot   *bool                          `json:"include_in_root,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *NestedProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NestedProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNestedProperty() *NestedProperty { _ = "STUB: not implemented"; return nil }

type NestedPropertyVariant interface {
	NestedPropertyCaster() *NestedProperty
}

func (s *NestedProperty) NestedPropertyCaster() *NestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *NestedProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
