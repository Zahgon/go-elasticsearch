package types

type FrequencyEncodingPreprocessor struct {
	FeatureName  string             `json:"feature_name"`
	Field        string             `json:"field"`
	FrequencyMap map[string]Float64 `json:"frequency_map"`
}

func (s *FrequencyEncodingPreprocessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFrequencyEncodingPreprocessor() *FrequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

type FrequencyEncodingPreprocessorVariant interface {
	FrequencyEncodingPreprocessorCaster() *FrequencyEncodingPreprocessor
}

func (s *FrequencyEncodingPreprocessor) FrequencyEncodingPreprocessorCaster() *FrequencyEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}
