package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

type BoxplotAggregation struct {
	Compression *Float64 `json:"compression,omitempty"`

	ExecutionHint *tdigestexecutionhint.TDigestExecutionHint `json:"execution_hint,omitempty"`

	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *BoxplotAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBoxplotAggregation() *BoxplotAggregation { _ = "STUB: not implemented"; return nil }

type BoxplotAggregationVariant interface {
	BoxplotAggregationCaster() *BoxplotAggregation
}

func (s *BoxplotAggregation) BoxplotAggregationCaster() *BoxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}
