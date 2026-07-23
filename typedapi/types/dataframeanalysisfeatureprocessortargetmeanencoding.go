package types

import (
	"encoding/json"
)

type DataframeAnalysisFeatureProcessorTargetMeanEncoding struct {
	DefaultValue int `json:"default_value"`

	FeatureName string `json:"feature_name"`

	Field string `json:"field"`

	TargetMap map[string]json.RawMessage `json:"target_map"`
}

func (s *DataframeAnalysisFeatureProcessorTargetMeanEncoding) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalysisFeatureProcessorTargetMeanEncoding() *DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisFeatureProcessorTargetMeanEncodingVariant interface {
	DataframeAnalysisFeatureProcessorTargetMeanEncodingCaster() *DataframeAnalysisFeatureProcessorTargetMeanEncoding
}

func (s *DataframeAnalysisFeatureProcessorTargetMeanEncoding) DataframeAnalysisFeatureProcessorTargetMeanEncodingCaster() *DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	_ = "STUB: not implemented"
	return nil
}
