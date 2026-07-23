package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorsimilarity"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type DenseVectorProperty struct {
	Dims    *int                           `json:"dims,omitempty"`
	Dynamic *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`

	ElementType *densevectorelementtype.DenseVectorElementType `json:"element_type,omitempty"`
	Fields      map[string]Property                            `json:"fields,omitempty"`
	IgnoreAbove *int                                           `json:"ignore_above,omitempty"`

	Index *bool `json:"index,omitempty"`

	IndexOptions *DenseVectorIndexOptions `json:"index_options,omitempty"`

	Meta       map[string]string   `json:"meta,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`

	Similarity          *densevectorsimilarity.DenseVectorSimilarity     `json:"similarity,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *DenseVectorProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DenseVectorProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDenseVectorProperty() *DenseVectorProperty { _ = "STUB: not implemented"; return nil }

type DenseVectorPropertyVariant interface {
	DenseVectorPropertyCaster() *DenseVectorProperty
}

func (s *DenseVectorProperty) DenseVectorPropertyCaster() *DenseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *DenseVectorProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
