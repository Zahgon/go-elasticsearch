package types

type CategorizeTextAggregation struct {
	CategorizationAnalyzer CategorizeTextAnalyzer `json:"categorization_analyzer,omitempty"`

	CategorizationFilters []string `json:"categorization_filters,omitempty"`

	Field string `json:"field"`

	MaxMatchedTokens *int `json:"max_matched_tokens,omitempty"`

	MaxUniqueTokens *int `json:"max_unique_tokens,omitempty"`

	MinDocCount *int `json:"min_doc_count,omitempty"`

	ShardMinDocCount *int `json:"shard_min_doc_count,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	SimilarityThreshold *int `json:"similarity_threshold,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *CategorizeTextAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCategorizeTextAggregation() *CategorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CategorizeTextAggregationVariant interface {
	CategorizeTextAggregationCaster() *CategorizeTextAggregation
}

func (s *CategorizeTextAggregation) CategorizeTextAggregationCaster() *CategorizeTextAggregation {
	_ = "STUB: not implemented"
	return nil
}
