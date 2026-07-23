package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type PercolatorProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *PercolatorProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PercolatorProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPercolatorProperty() *PercolatorProperty { _ = "STUB: not implemented"; return nil }

type PercolatorPropertyVariant interface {
	PercolatorPropertyCaster() *PercolatorProperty
}

func (s *PercolatorProperty) PercolatorPropertyCaster() *PercolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *PercolatorProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
