package types

type DataframeAnalysisFeatureProcessor struct {
	FrequencyEncoding *DataframeAnalysisFeatureProcessorFrequencyEncoding `json:"frequency_encoding,omitempty"`

	MultiEncoding *DataframeAnalysisFeatureProcessorMultiEncoding `json:"multi_encoding,omitempty"`

	NGramEncoding *DataframeAnalysisFeatureProcessorNGramEncoding `json:"n_gram_encoding,omitempty"`

	OneHotEncoding *DataframeAnalysisFeatureProcessorOneHotEncoding `json:"one_hot_encoding,omitempty"`

	TargetMeanEncoding *DataframeAnalysisFeatureProcessorTargetMeanEncoding `json:"target_mean_encoding,omitempty"`
}

func NewDataframeAnalysisFeatureProcessor() *DataframeAnalysisFeatureProcessor {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisFeatureProcessorVariant interface {
	DataframeAnalysisFeatureProcessorCaster() *DataframeAnalysisFeatureProcessor
}

func (s *DataframeAnalysisFeatureProcessor) DataframeAnalysisFeatureProcessorCaster() *DataframeAnalysisFeatureProcessor {
	_ = "STUB: not implemented"
	return nil
}
