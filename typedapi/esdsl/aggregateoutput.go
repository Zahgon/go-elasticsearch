package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _aggregateOutput struct {
	v *types.AggregateOutput
}

func NewAggregateOutput() *_aggregateOutput { _ = "STUB: not implemented"; return nil }

func (s *_aggregateOutput) Exponent(exponent types.WeightsVariant) *_aggregateOutput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateOutput) LogisticRegression(logisticregression types.WeightsVariant) *_aggregateOutput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateOutput) WeightedMode(weightedmode types.WeightsVariant) *_aggregateOutput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateOutput) WeightedSum(weightedsum types.WeightsVariant) *_aggregateOutput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateOutput) AggregateOutputCaster() *types.AggregateOutput {
	_ = "STUB: not implemented"
	return nil
}
