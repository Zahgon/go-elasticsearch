package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type PointProperty struct {
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           *string                                          `json:"null_value,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *PointProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s PointProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewPointProperty() *PointProperty { _ = "STUB: not implemented"; return nil }

type PointPropertyVariant interface {
	PointPropertyCaster() *PointProperty
}

func (s *PointProperty) PointPropertyCaster() *PointProperty { _ = "STUB: not implemented"; return nil }

func (s *PointProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
