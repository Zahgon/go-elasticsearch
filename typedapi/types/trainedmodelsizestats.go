package types

type TrainedModelSizeStats struct {
	ModelSizeBytes ByteSize `json:"model_size_bytes"`

	RequiredNativeMemoryBytes ByteSize `json:"required_native_memory_bytes"`
}

func (s *TrainedModelSizeStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelSizeStats() *TrainedModelSizeStats { _ = "STUB: not implemented"; return nil }
