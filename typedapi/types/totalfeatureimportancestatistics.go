package types

type TotalFeatureImportanceStatistics struct {
	Max int `json:"max"`

	MeanMagnitude Float64 `json:"mean_magnitude"`

	Min int `json:"min"`
}

func (s *TotalFeatureImportanceStatistics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTotalFeatureImportanceStatistics() *TotalFeatureImportanceStatistics {
	_ = "STUB: not implemented"
	return nil
}
