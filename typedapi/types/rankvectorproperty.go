package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rankvectorelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type RankVectorProperty struct {
	Dims        *int                                         `json:"dims,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping               `json:"dynamic,omitempty"`
	ElementType *rankvectorelementtype.RankVectorElementType `json:"element_type,omitempty"`
	Fields      map[string]Property                          `json:"fields,omitempty"`
	IgnoreAbove *int                                         `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *RankVectorProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s RankVectorProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRankVectorProperty() *RankVectorProperty { _ = "STUB: not implemented"; return nil }

type RankVectorPropertyVariant interface {
	RankVectorPropertyCaster() *RankVectorProperty
}

func (s *RankVectorProperty) RankVectorPropertyCaster() *RankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *RankVectorProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
