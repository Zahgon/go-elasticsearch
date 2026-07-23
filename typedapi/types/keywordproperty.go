package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type KeywordProperty struct {
	Boost               *Float64                       `json:"boost,omitempty"`
	CopyTo              []string                       `json:"copy_to,omitempty"`
	DocValues           *bool                          `json:"doc_values,omitempty"`
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool                          `json:"eager_global_ordinals,omitempty"`
	Fields              map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`
	Index               *bool                          `json:"index,omitempty"`
	IndexOptions        *indexoptions.IndexOptions     `json:"index_options,omitempty"`

	Meta                     map[string]string                                `json:"meta,omitempty"`
	Normalizer               *string                                          `json:"normalizer,omitempty"`
	Norms                    *bool                                            `json:"norms,omitempty"`
	NullValue                *string                                          `json:"null_value,omitempty"`
	OnScriptError            *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	Properties               map[string]Property                              `json:"properties,omitempty"`
	Script                   *Script                                          `json:"script,omitempty"`
	Similarity               *string                                          `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace *bool                                            `json:"split_queries_on_whitespace,omitempty"`
	Store                    *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep      *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`

	TimeSeriesDimension *bool  `json:"time_series_dimension,omitempty"`
	Type                string `json:"type,omitempty"`
}

func (s *KeywordProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s KeywordProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewKeywordProperty() *KeywordProperty { _ = "STUB: not implemented"; return nil }

type KeywordPropertyVariant interface {
	KeywordPropertyCaster() *KeywordProperty
}

func (s *KeywordProperty) KeywordPropertyCaster() *KeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeywordProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
