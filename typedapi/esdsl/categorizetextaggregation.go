package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _categorizeTextAggregation struct {
	v *types.CategorizeTextAggregation
}

func NewCategorizeTextAggregation() *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) CategorizationAnalyzer(categorizetextanalyzer types.CategorizeTextAnalyzerVariant) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) CategorizationFilters(categorizationfilters ...string) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) Field(field string) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) MaxMatchedTokens(maxmatchedtokens int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) MaxUniqueTokens(maxuniquetokens int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) MinDocCount(mindoccount int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) ShardMinDocCount(shardmindoccount int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) ShardSize(shardsize int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) SimilarityThreshold(similaritythreshold int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) Size(size int) *_categorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizeTextAggregation) CategorizeTextAggregationCaster() *types.CategorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}
