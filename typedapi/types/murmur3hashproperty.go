package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type Murmur3HashProperty struct {
	CopyTo      []string                       `json:"copy_to,omitempty"`
	DocValues   *bool                          `json:"doc_values,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *Murmur3HashProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s Murmur3HashProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMurmur3HashProperty() *Murmur3HashProperty { _ = "STUB: not implemented"; return nil }

type Murmur3HashPropertyVariant interface {
	Murmur3HashPropertyCaster() *Murmur3HashProperty
}

func (s *Murmur3HashProperty) Murmur3HashPropertyCaster() *Murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *Murmur3HashProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
