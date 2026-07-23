package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _oneHotEncodingPreprocessor struct {
	v *types.OneHotEncodingPreprocessor
}

func NewOneHotEncodingPreprocessor(field string) *_oneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_oneHotEncodingPreprocessor) Field(field string) *_oneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_oneHotEncodingPreprocessor) HotMap(hotmap map[string]string) *_oneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_oneHotEncodingPreprocessor) AddHotMap(key string, value string) *_oneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_oneHotEncodingPreprocessor) PreprocessorCaster() *types.Preprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_oneHotEncodingPreprocessor) OneHotEncodingPreprocessorCaster() *types.OneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}
