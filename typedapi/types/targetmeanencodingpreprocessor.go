package types

type TargetMeanEncodingPreprocessor struct {
	DefaultValue Float64            `json:"default_value"`
	FeatureName  string             `json:"feature_name"`
	Field        string             `json:"field"`
	TargetMap    map[string]Float64 `json:"target_map"`
}

func (s *TargetMeanEncodingPreprocessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTargetMeanEncodingPreprocessor() *TargetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

type TargetMeanEncodingPreprocessorVariant interface {
	TargetMeanEncodingPreprocessorCaster() *TargetMeanEncodingPreprocessor
}

func (s *TargetMeanEncodingPreprocessor) TargetMeanEncodingPreprocessorCaster() *TargetMeanEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}
