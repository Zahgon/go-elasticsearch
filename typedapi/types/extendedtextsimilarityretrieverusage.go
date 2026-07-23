package types

type ExtendedTextSimilarityRetrieverUsage struct {
	ChunkRescorer *int64 `json:"chunk_rescorer,omitempty"`
}

func (s *ExtendedTextSimilarityRetrieverUsage) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedTextSimilarityRetrieverUsage() *ExtendedTextSimilarityRetrieverUsage {
	_ = "STUB: not implemented"
	return nil
}
