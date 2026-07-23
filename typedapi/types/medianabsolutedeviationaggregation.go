package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

type MedianAbsoluteDeviationAggregation struct {
	Compression *Float64 `json:"compression,omitempty"`

	ExecutionHint *tdigestexecutionhint.TDigestExecutionHint `json:"execution_hint,omitempty"`

	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *MedianAbsoluteDeviationAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMedianAbsoluteDeviationAggregation() *MedianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

type MedianAbsoluteDeviationAggregationVariant interface {
	MedianAbsoluteDeviationAggregationCaster() *MedianAbsoluteDeviationAggregation
}

func (s *MedianAbsoluteDeviationAggregation) MedianAbsoluteDeviationAggregationCaster() *MedianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}
