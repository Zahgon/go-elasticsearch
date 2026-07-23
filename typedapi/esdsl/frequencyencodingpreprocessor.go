package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _frequencyEncodingPreprocessor struct {
	v *types.FrequencyEncodingPreprocessor
}

func NewFrequencyEncodingPreprocessor(featurename string, field string) *_frequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequencyEncodingPreprocessor) FeatureName(featurename string) *_frequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequencyEncodingPreprocessor) Field(field string) *_frequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequencyEncodingPreprocessor) FrequencyMap(frequencymap map[string]types.Float64) *_frequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequencyEncodingPreprocessor) AddFrequencyMap(key string, value types.Float64) *_frequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequencyEncodingPreprocessor) PreprocessorCaster() *types.Preprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequencyEncodingPreprocessor) FrequencyEncodingPreprocessorCaster() *types.FrequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}
