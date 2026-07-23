package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rareTermsAggregation struct {
	v *types.RareTermsAggregation
}

func NewRareTermsAggregation() *_rareTermsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_rareTermsAggregation) Exclude(termsexcludes ...string) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) Field(field string) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) Include(termsinclude types.TermsIncludeVariant) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) MaxDocCount(maxdoccount int64) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) Missing(missing types.MissingVariant) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) Precision(precision types.Float64) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) ValueType(valuetype string) *_rareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rareTermsAggregation) RareTermsAggregationCaster() *types.RareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
