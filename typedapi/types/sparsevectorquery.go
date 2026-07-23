package types

type SparseVectorQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	InferenceId *string `json:"inference_id,omitempty"`

	Prune *bool `json:"prune,omitempty"`

	PruningConfig *TokenPruningConfig `json:"pruning_config,omitempty"`

	Query      *string `json:"query,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	QueryVector map[string]float32 `json:"query_vector,omitempty"`
}

func (s *SparseVectorQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSparseVectorQuery() *SparseVectorQuery { _ = "STUB: not implemented"; return nil }

type SparseVectorQueryVariant interface {
	SparseVectorQueryCaster() *SparseVectorQuery
}

func (s *SparseVectorQuery) SparseVectorQueryCaster() *SparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}
