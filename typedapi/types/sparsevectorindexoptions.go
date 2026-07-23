package types

type SparseVectorIndexOptions struct {
	Prune *bool `json:"prune,omitempty"`

	PruningConfig *TokenPruningConfig `json:"pruning_config,omitempty"`
}

func (s *SparseVectorIndexOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSparseVectorIndexOptions() *SparseVectorIndexOptions { _ = "STUB: not implemented"; return nil }

type SparseVectorIndexOptionsVariant interface {
	SparseVectorIndexOptionsCaster() *SparseVectorIndexOptions
}

func (s *SparseVectorIndexOptions) SparseVectorIndexOptionsCaster() *SparseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}
