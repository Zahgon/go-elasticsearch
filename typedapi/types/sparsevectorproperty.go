package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type SparseVectorProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	IndexOptions *SparseVectorIndexOptions `json:"index_options,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *SparseVectorProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SparseVectorProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSparseVectorProperty() *SparseVectorProperty { _ = "STUB: not implemented"; return nil }

type SparseVectorPropertyVariant interface {
	SparseVectorPropertyCaster() *SparseVectorProperty
}

func (s *SparseVectorProperty) SparseVectorPropertyCaster() *SparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *SparseVectorProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
