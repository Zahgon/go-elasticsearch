package types

type DataframeAnalysisFeatureProcessorOneHotEncoding struct {
	Field string `json:"field"`

	HotMap string `json:"hot_map"`
}

func (s *DataframeAnalysisFeatureProcessorOneHotEncoding) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalysisFeatureProcessorOneHotEncoding() *DataframeAnalysisFeatureProcessorOneHotEncoding {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisFeatureProcessorOneHotEncodingVariant interface {
	DataframeAnalysisFeatureProcessorOneHotEncodingCaster() *DataframeAnalysisFeatureProcessorOneHotEncoding
}

func (s *DataframeAnalysisFeatureProcessorOneHotEncoding) DataframeAnalysisFeatureProcessorOneHotEncodingCaster() *DataframeAnalysisFeatureProcessorOneHotEncoding {
	_ = "STUB: not implemented"
	return nil
}
