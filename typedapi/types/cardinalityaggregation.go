package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cardinalityexecutionmode"
)

type CardinalityAggregation struct {
	ExecutionHint *cardinalityexecutionmode.CardinalityExecutionMode `json:"execution_hint,omitempty"`

	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	PrecisionThreshold *int    `json:"precision_threshold,omitempty"`
	Rehash             *bool   `json:"rehash,omitempty"`
	Script             *Script `json:"script,omitempty"`
}

func (s *CardinalityAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCardinalityAggregation() *CardinalityAggregation { _ = "STUB: not implemented"; return nil }

type CardinalityAggregationVariant interface {
	CardinalityAggregationCaster() *CardinalityAggregation
}

func (s *CardinalityAggregation) CardinalityAggregationCaster() *CardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}
