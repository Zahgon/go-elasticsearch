package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _targetMeanEncodingPreprocessor struct {
	v *types.TargetMeanEncodingPreprocessor
}

func NewTargetMeanEncodingPreprocessor(defaultvalue types.Float64, featurename string, field string) *_targetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) DefaultValue(defaultvalue types.Float64) *_targetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) FeatureName(featurename string) *_targetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) Field(field string) *_targetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) TargetMap(targetmap map[string]types.Float64) *_targetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) AddTargetMap(key string, value types.Float64) *_targetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) PreprocessorCaster() *types.Preprocessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_targetMeanEncodingPreprocessor) TargetMeanEncodingPreprocessorCaster() *types.TargetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}
