package types

type DataframeAnalysisFeatureProcessorFrequencyEncoding struct {
	FeatureName string `json:"feature_name"`
	Field       string `json:"field"`

	FrequencyMap map[string]Float64 `json:"frequency_map"`
}

func (s *DataframeAnalysisFeatureProcessorFrequencyEncoding) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalysisFeatureProcessorFrequencyEncoding() *DataframeAnalysisFeatureProcessorFrequencyEncoding {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisFeatureProcessorFrequencyEncodingVariant interface {
	DataframeAnalysisFeatureProcessorFrequencyEncodingCaster() *DataframeAnalysisFeatureProcessorFrequencyEncoding
}

func (s *DataframeAnalysisFeatureProcessorFrequencyEncoding) DataframeAnalysisFeatureProcessorFrequencyEncodingCaster() *DataframeAnalysisFeatureProcessorFrequencyEncoding {
	_ = "STUB: not implemented"
	return nil
}
