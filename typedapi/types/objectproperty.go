package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type ObjectProperty struct {
	CopyTo      []string                       `json:"copy_to,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled     *bool                          `json:"enabled,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	Subobjects          *subobjects.Subobjects                           `json:"subobjects,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *ObjectProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ObjectProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewObjectProperty() *ObjectProperty { _ = "STUB: not implemented"; return nil }

type ObjectPropertyVariant interface {
	ObjectPropertyCaster() *ObjectProperty
}

func (s *ObjectProperty) ObjectPropertyCaster() *ObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *ObjectProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
