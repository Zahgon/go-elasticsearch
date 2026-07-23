package types

type SemanticTextIndexOptions struct {
	DenseVector  *DenseVectorIndexOptions  `json:"dense_vector,omitempty"`
	SparseVector *SparseVectorIndexOptions `json:"sparse_vector,omitempty"`
}

func NewSemanticTextIndexOptions() *SemanticTextIndexOptions { _ = "STUB: not implemented"; return nil }

type SemanticTextIndexOptionsVariant interface {
	SemanticTextIndexOptionsCaster() *SemanticTextIndexOptions
}

func (s *SemanticTextIndexOptions) SemanticTextIndexOptionsCaster() *SemanticTextIndexOptions {
	_ = "STUB: not implemented"
	return nil
}
