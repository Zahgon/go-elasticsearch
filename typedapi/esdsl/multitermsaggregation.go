package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationcollectmode"
)

type _multiTermsAggregation struct {
	v *types.MultiTermsAggregation
}

func NewMultiTermsAggregation() *_multiTermsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_multiTermsAggregation) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) MinDocCount(mindoccount int64) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) Order(aggregateorder types.AggregateOrderVariant) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) ShardMinDocCount(shardmindoccount int64) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) ShardSize(shardsize int) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) ShowTermDocCountError(showtermdoccounterror bool) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) Size(size int) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) Terms(terms ...types.MultiTermLookupVariant) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) TermsValues(termsvalues []types.MultiTermLookup) *_multiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermsAggregation) MultiTermsAggregationCaster() *types.MultiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
