package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _categorizationAnalyzerDefinition struct {
	v *types.CategorizationAnalyzerDefinition
}

func NewCategorizationAnalyzerDefinition() *_categorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizationAnalyzerDefinition) CharFilter(charfilters ...types.CharFilterVariant) *_categorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizationAnalyzerDefinition) CharFilterValues(charfiltervalues []types.CharFilter) *_categorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizationAnalyzerDefinition) Filter(filters ...types.TokenFilterVariant) *_categorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizationAnalyzerDefinition) FilterValues(filtervalues []types.TokenFilter) *_categorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizationAnalyzerDefinition) Tokenizer(tokenizer types.TokenizerVariant) *_categorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_categorizationAnalyzerDefinition) CategorizationAnalyzerDefinitionCaster() *types.CategorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}
