package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
)

type TextProperty struct {
	Analyzer                 *string                        `json:"analyzer,omitempty"`
	Boost                    *Float64                       `json:"boost,omitempty"`
	CopyTo                   []string                       `json:"copy_to,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                          `json:"eager_global_ordinals,omitempty"`
	Fielddata                *bool                          `json:"fielddata,omitempty"`
	FielddataFrequencyFilter *FielddataFrequencyFilter      `json:"fielddata_frequency_filter,omitempty"`
	Fields                   map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	IndexPhrases             *bool                          `json:"index_phrases,omitempty"`
	IndexPrefixes            *TextIndexPrefixes             `json:"index_prefixes,omitempty"`

	Meta                 map[string]string                                `json:"meta,omitempty"`
	Norms                *bool                                            `json:"norms,omitempty"`
	PositionIncrementGap *int                                             `json:"position_increment_gap,omitempty"`
	Properties           map[string]Property                              `json:"properties,omitempty"`
	SearchAnalyzer       *string                                          `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  *string                                          `json:"search_quote_analyzer,omitempty"`
	Similarity           *string                                          `json:"similarity,omitempty"`
	Store                *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep  *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TermVector           *termvectoroption.TermVectorOption               `json:"term_vector,omitempty"`
	Type                 string                                           `json:"type,omitempty"`
}

func (s *TextProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s TextProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTextProperty() *TextProperty { _ = "STUB: not implemented"; return nil }

type TextPropertyVariant interface {
	TextPropertyCaster() *TextProperty
}

func (s *TextProperty) TextPropertyCaster() *TextProperty { _ = "STUB: not implemented"; return nil }

func (s *TextProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
