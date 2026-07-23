package types

type OneHotEncodingPreprocessor struct {
	Field  string            `json:"field"`
	HotMap map[string]string `json:"hot_map"`
}

func (s *OneHotEncodingPreprocessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOneHotEncodingPreprocessor() *OneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

type OneHotEncodingPreprocessorVariant interface {
	OneHotEncodingPreprocessorCaster() *OneHotEncodingPreprocessor
}

func (s *OneHotEncodingPreprocessor) OneHotEncodingPreprocessorCaster() *OneHotEncodingPreprocessor {
	_ = "STUB: not implemented"
	return nil
}
