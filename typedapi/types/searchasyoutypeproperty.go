package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
)

type SearchAsYouTypeProperty struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	CopyTo         []string                       `json:"copy_to,omitempty"`
	Dynamic        *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields         map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove    *int                           `json:"ignore_above,omitempty"`
	Index          *bool                          `json:"index,omitempty"`
	IndexOptions   *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	MaxShingleSize *int                           `json:"max_shingle_size,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Norms               *bool                                            `json:"norms,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SearchAnalyzer      *string                                          `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer *string                                          `json:"search_quote_analyzer,omitempty"`
	Similarity          *string                                          `json:"similarity,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TermVector          *termvectoroption.TermVectorOption               `json:"term_vector,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *SearchAsYouTypeProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SearchAsYouTypeProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSearchAsYouTypeProperty() *SearchAsYouTypeProperty { _ = "STUB: not implemented"; return nil }

type SearchAsYouTypePropertyVariant interface {
	SearchAsYouTypePropertyCaster() *SearchAsYouTypeProperty
}

func (s *SearchAsYouTypeProperty) SearchAsYouTypePropertyCaster() *SearchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *SearchAsYouTypeProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
