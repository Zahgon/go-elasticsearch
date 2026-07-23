package types

type TrainedModelInferenceStats struct {
	CacheMissCount int `json:"cache_miss_count"`

	FailureCount int `json:"failure_count"`

	InferenceCount int `json:"inference_count"`

	MissingAllFieldsCount int `json:"missing_all_fields_count"`

	Timestamp int64 `json:"timestamp"`
}

func (s *TrainedModelInferenceStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelInferenceStats() *TrainedModelInferenceStats {
	_ = "STUB: not implemented"
	return nil
}
