package types

import (
	"encoding/json"
)

type ApiKeyAggregationContainer struct {
	AdditionalApiKeyAggregationContainerProperty map[string]json.RawMessage `json:"-"`

	Aggregations map[string]ApiKeyAggregationContainer `json:"aggregations,omitempty"`

	Cardinality *CardinalityAggregation `json:"cardinality,omitempty"`

	Composite *CompositeAggregation `json:"composite,omitempty"`

	DateRange *DateRangeAggregation `json:"date_range,omitempty"`

	Filter *ApiKeyQueryContainer `json:"filter,omitempty"`

	Filters *ApiKeyFiltersAggregation `json:"filters,omitempty"`
	Meta    Metadata                  `json:"meta,omitempty"`
	Missing *MissingAggregation       `json:"missing,omitempty"`

	Range *RangeAggregation `json:"range,omitempty"`

	Terms *TermsAggregation `json:"terms,omitempty"`

	ValueCount *ValueCountAggregation `json:"value_count,omitempty"`
}

func (s *ApiKeyAggregationContainer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ApiKeyAggregationContainer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewApiKeyAggregationContainer() *ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

type ApiKeyAggregationContainerVariant interface {
	ApiKeyAggregationContainerCaster() *ApiKeyAggregationContainer
}

func (s *ApiKeyAggregationContainer) ApiKeyAggregationContainerCaster() *ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}
