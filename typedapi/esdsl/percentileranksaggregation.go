package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _percentileRanksAggregation struct {
	v *types.PercentileRanksAggregation
}

func NewPercentileRanksAggregation() *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Hdr(hdr types.HdrMethodVariant) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Keyed(keyed bool) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Tdigest(tdigest types.TDigestVariant) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Values(values []types.Float64) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Field(field string) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Format(format string) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Missing(missing types.MissingVariant) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) Script(script types.ScriptVariant) *_percentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentileRanksAggregation) PercentileRanksAggregationCaster() *types.PercentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}
