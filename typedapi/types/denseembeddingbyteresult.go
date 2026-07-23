package types

type DenseEmbeddingByteResult struct {
	Embedding []byte `json:"embedding"`
}

func (s *DenseEmbeddingByteResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDenseEmbeddingByteResult() *DenseEmbeddingByteResult { _ = "STUB: not implemented"; return nil }
