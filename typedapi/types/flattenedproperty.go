package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type FlattenedProperty struct {
	Boost               *Float64                       `json:"boost,omitempty"`
	DepthLimit          *int                           `json:"depth_limit,omitempty"`
	DocValues           *bool                          `json:"doc_values,omitempty"`
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool                          `json:"eager_global_ordinals,omitempty"`
	Fields              map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`
	Index               *bool                          `json:"index,omitempty"`
	IndexOptions        *indexoptions.IndexOptions     `json:"index_options,omitempty"`

	Meta                     map[string]string                                `json:"meta,omitempty"`
	NullValue                *string                                          `json:"null_value,omitempty"`
	Properties               map[string]Property                              `json:"properties,omitempty"`
	Similarity               *string                                          `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace *bool                                            `json:"split_queries_on_whitespace,omitempty"`
	SyntheticSourceKeep      *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesDimensions     []string                                         `json:"time_series_dimensions,omitempty"`
	Type                     string                                           `json:"type,omitempty"`
}

func (s *FlattenedProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s FlattenedProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFlattenedProperty() *FlattenedProperty { _ = "STUB: not implemented"; return nil }

type FlattenedPropertyVariant interface {
	FlattenedPropertyCaster() *FlattenedProperty
}

func (s *FlattenedProperty) FlattenedPropertyCaster() *FlattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *FlattenedProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
