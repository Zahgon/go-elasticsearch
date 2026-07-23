package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingsAnalysis struct {
	v *types.IndexSettingsAnalysis
}

func NewIndexSettingsAnalysis() *_indexSettingsAnalysis { _ = "STUB: not implemented"; return nil }

func (s *_indexSettingsAnalysis) Analyzer(analyzer map[string]types.Analyzer) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) AddAnalyzer(key string, value types.AnalyzerVariant) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) CharFilter(charfilter map[string]types.CharFilter) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) AddCharFilter(key string, value types.CharFilterVariant) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) Filter(filter map[string]types.TokenFilter) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) AddFilter(key string, value types.TokenFilterVariant) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) Normalizer(normalizer map[string]types.Normalizer) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) AddNormalizer(key string, value types.NormalizerVariant) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) Tokenizer(tokenizer map[string]types.Tokenizer) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) AddTokenizer(key string, value types.TokenizerVariant) *_indexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsAnalysis) IndexSettingsAnalysisCaster() *types.IndexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}
