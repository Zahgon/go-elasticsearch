package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ipPrefixAggregation struct {
	v *types.IpPrefixAggregation
}

func NewIpPrefixAggregation(prefixlength int) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) AppendPrefixLength(appendprefixlength bool) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) Field(field string) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) IsIpv6(isipv6 bool) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) Keyed(keyed bool) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) MinDocCount(mindoccount int64) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) PrefixLength(prefixlength int) *_ipPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipPrefixAggregation) IpPrefixAggregationCaster() *types.IpPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}
