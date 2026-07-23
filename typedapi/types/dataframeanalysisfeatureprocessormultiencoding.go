package types

type DataframeAnalysisFeatureProcessorMultiEncoding struct {
	Processors []int `json:"processors"`
}

func NewDataframeAnalysisFeatureProcessorMultiEncoding() *DataframeAnalysisFeatureProcessorMultiEncoding {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisFeatureProcessorMultiEncodingVariant interface {
	DataframeAnalysisFeatureProcessorMultiEncodingCaster() *DataframeAnalysisFeatureProcessorMultiEncoding
}

func (s *DataframeAnalysisFeatureProcessorMultiEncoding) DataframeAnalysisFeatureProcessorMultiEncodingCaster() *DataframeAnalysisFeatureProcessorMultiEncoding {
	_ = "STUB: not implemented"
	return nil
}
