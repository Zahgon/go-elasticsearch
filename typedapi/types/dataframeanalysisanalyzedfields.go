package types

type DataframeAnalysisAnalyzedFields struct {
	Excludes []string `json:"excludes,omitempty"`

	Includes []string `json:"includes,omitempty"`
}

func (s *DataframeAnalysisAnalyzedFields) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalysisAnalyzedFields() *DataframeAnalysisAnalyzedFields {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisAnalyzedFieldsVariant interface {
	DataframeAnalysisAnalyzedFieldsCaster() *DataframeAnalysisAnalyzedFields
}

func (s *DataframeAnalysisAnalyzedFields) DataframeAnalysisAnalyzedFieldsCaster() *DataframeAnalysisAnalyzedFields {
	_ = "STUB: not implemented"
	return nil
}
