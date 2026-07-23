package types

type TotalFeatureImportanceClass struct {
	ClassName string `json:"class_name"`

	Importance []TotalFeatureImportanceStatistics `json:"importance"`
}

func (s *TotalFeatureImportanceClass) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTotalFeatureImportanceClass() *TotalFeatureImportanceClass {
	_ = "STUB: not implemented"
	return nil
}
