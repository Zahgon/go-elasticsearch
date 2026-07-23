package types

type SparseEmbeddingResult struct {
	Embedding SparseVector `json:"embedding"`

	IsTruncated bool `json:"is_truncated"`
}

func (s *SparseEmbeddingResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSparseEmbeddingResult() *SparseEmbeddingResult { _ = "STUB: not implemented"; return nil }
