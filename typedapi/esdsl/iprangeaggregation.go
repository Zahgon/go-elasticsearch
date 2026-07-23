package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ipRangeAggregation struct {
	v *types.IpRangeAggregation
}

func NewIpRangeAggregation() *_ipRangeAggregation { _ = "STUB: not implemented"; return nil }

func (s *_ipRangeAggregation) Field(field string) *_ipRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeAggregation) Ranges(ranges ...types.IpRangeAggregationRangeVariant) *_ipRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeAggregation) RangesValues(rangesvalues []types.IpRangeAggregationRange) *_ipRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeAggregation) IpRangeAggregationCaster() *types.IpRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}
