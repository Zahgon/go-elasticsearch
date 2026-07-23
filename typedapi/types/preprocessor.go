package types

type Preprocessor struct {
	FrequencyEncoding  *FrequencyEncodingPreprocessor  `json:"frequency_encoding,omitempty"`
	OneHotEncoding     *OneHotEncodingPreprocessor     `json:"one_hot_encoding,omitempty"`
	TargetMeanEncoding *TargetMeanEncodingPreprocessor `json:"target_mean_encoding,omitempty"`
}

func NewPreprocessor() *Preprocessor { _ = "STUB: not implemented"; return nil }

type PreprocessorVariant interface {
	PreprocessorCaster() *Preprocessor
}

func (s *Preprocessor) PreprocessorCaster() *Preprocessor { _ = "STUB: not implemented"; return nil }
