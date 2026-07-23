package types

type MlInferenceIngestProcessorCount struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	Sum int64 `json:"sum"`
}

func (s *MlInferenceIngestProcessorCount) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMlInferenceIngestProcessorCount() *MlInferenceIngestProcessorCount {
	_ = "STUB: not implemented"
	return nil
}
