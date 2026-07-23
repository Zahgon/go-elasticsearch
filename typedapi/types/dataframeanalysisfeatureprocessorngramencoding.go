package types

type DataframeAnalysisFeatureProcessorNGramEncoding struct {
	Custom *bool `json:"custom,omitempty"`

	FeaturePrefix *string `json:"feature_prefix,omitempty"`

	Field string `json:"field"`

	Length *int `json:"length,omitempty"`

	NGrams []int `json:"n_grams"`

	Start *int `json:"start,omitempty"`
}

func (s *DataframeAnalysisFeatureProcessorNGramEncoding) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalysisFeatureProcessorNGramEncoding() *DataframeAnalysisFeatureProcessorNGramEncoding {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisFeatureProcessorNGramEncodingVariant interface {
	DataframeAnalysisFeatureProcessorNGramEncodingCaster() *DataframeAnalysisFeatureProcessorNGramEncoding
}

func (s *DataframeAnalysisFeatureProcessorNGramEncoding) DataframeAnalysisFeatureProcessorNGramEncodingCaster() *DataframeAnalysisFeatureProcessorNGramEncoding {
	_ = "STUB: not implemented"
	return nil
}
