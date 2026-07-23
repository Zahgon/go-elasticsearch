package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _percentilesAggregation struct {
	v *types.PercentilesAggregation
}

func NewPercentilesAggregation() *_percentilesAggregation { _ = "STUB: not implemented"; return nil }

func (s *_percentilesAggregation) Hdr(hdr types.HdrMethodVariant) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Keyed(keyed bool) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Percents(percents ...types.Float64) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Tdigest(tdigest types.TDigestVariant) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Field(field string) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Format(format string) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Missing(missing types.MissingVariant) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) Script(script types.ScriptVariant) *_percentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesAggregation) PercentilesAggregationCaster() *types.PercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}
