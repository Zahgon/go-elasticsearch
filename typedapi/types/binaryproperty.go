package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type BinaryProperty struct {
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

func (s *BinaryProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s BinaryProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewBinaryProperty() *BinaryProperty { _ = "STUB: not implemented"; return nil }

type BinaryPropertyVariant interface {
	BinaryPropertyCaster() *BinaryProperty
}

func (s *BinaryProperty) BinaryPropertyCaster() *BinaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *BinaryProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
