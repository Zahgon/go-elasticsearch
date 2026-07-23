package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type DynamicProperty struct {
	Analyzer            *string                        `json:"analyzer,omitempty"`
	Boost               *Float64                       `json:"boost,omitempty"`
	Coerce              *bool                          `json:"coerce,omitempty"`
	CopyTo              []string                       `json:"copy_to,omitempty"`
	DocValues           *bool                          `json:"doc_values,omitempty"`
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool                          `json:"eager_global_ordinals,omitempty"`
	Enabled             *bool                          `json:"enabled,omitempty"`
	Fields              map[string]Property            `json:"fields,omitempty"`
	Format              *string                        `json:"format,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed     *bool                          `json:"ignore_malformed,omitempty"`
	Index               *bool                          `json:"index,omitempty"`
	IndexOptions        *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	IndexPhrases        *bool                          `json:"index_phrases,omitempty"`
	IndexPrefixes       *TextIndexPrefixes             `json:"index_prefixes,omitempty"`
	Locale              *string                        `json:"locale,omitempty"`

	Meta                 map[string]string                                `json:"meta,omitempty"`
	Norms                *bool                                            `json:"norms,omitempty"`
	NullValue            FieldValue                                       `json:"null_value,omitempty"`
	OnScriptError        *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	PositionIncrementGap *int                                             `json:"position_increment_gap,omitempty"`
	PrecisionStep        *int                                             `json:"precision_step,omitempty"`
	Properties           map[string]Property                              `json:"properties,omitempty"`
	Script               *Script                                          `json:"script,omitempty"`
	SearchAnalyzer       *string                                          `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  *string                                          `json:"search_quote_analyzer,omitempty"`
	Store                *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep  *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TermVector           *termvectoroption.TermVectorOption               `json:"term_vector,omitempty"`
	TimeSeriesMetric     *timeseriesmetrictype.TimeSeriesMetricType       `json:"time_series_metric,omitempty"`
	Type                 string                                           `json:"type,omitempty"`
}

func (s *DynamicProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s DynamicProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDynamicProperty() *DynamicProperty { _ = "STUB: not implemented"; return nil }

type DynamicPropertyVariant interface {
	DynamicPropertyCaster() *DynamicProperty
}

func (s *DynamicProperty) DynamicPropertyCaster() *DynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *DynamicProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
