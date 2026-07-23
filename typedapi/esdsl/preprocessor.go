package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _preprocessor struct {
	v *types.Preprocessor
}

func NewPreprocessor() *_preprocessor { _ = "STUB: not implemented"; return nil }

func (s *_preprocessor) FrequencyEncoding(frequencyencoding types.FrequencyEncodingPreprocessorVariant) *_preprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_preprocessor) OneHotEncoding(onehotencoding types.OneHotEncodingPreprocessorVariant) *_preprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_preprocessor) TargetMeanEncoding(targetmeanencoding types.TargetMeanEncodingPreprocessorVariant) *_preprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_preprocessor) PreprocessorCaster() *types.Preprocessor {
	_ = "STUB: not implemented"
	return nil
}
