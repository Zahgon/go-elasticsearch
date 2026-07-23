package types

type DenseEmbeddingResult struct {
	Embedding []float32 `json:"embedding"`
}

func (s *DenseEmbeddingResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDenseEmbeddingResult() *DenseEmbeddingResult { _ = "STUB: not implemented"; return nil }
