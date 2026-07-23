package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sampleraggregationexecutionhint"
)

type DiversifiedSamplerAggregation struct {
	ExecutionHint *sampleraggregationexecutionhint.SamplerAggregationExecutionHint `json:"execution_hint,omitempty"`

	Field *string `json:"field,omitempty"`

	MaxDocsPerValue *int    `json:"max_docs_per_value,omitempty"`
	Script          *Script `json:"script,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`
}

func (s *DiversifiedSamplerAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDiversifiedSamplerAggregation() *DiversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

type DiversifiedSamplerAggregationVariant interface {
	DiversifiedSamplerAggregationCaster() *DiversifiedSamplerAggregation
}

func (s *DiversifiedSamplerAggregation) DiversifiedSamplerAggregationCaster() *DiversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}
