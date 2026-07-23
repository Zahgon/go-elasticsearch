package types

type KnnSearchQuery struct {
	Field string `json:"field"`

	K int `json:"k"`

	NumCandidates int `json:"num_candidates"`

	QueryVector []float32 `json:"query_vector"`
}

func (s *KnnSearchQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewKnnSearchQuery() *KnnSearchQuery { _ = "STUB: not implemented"; return nil }

type KnnSearchQueryVariant interface {
	KnnSearchQueryCaster() *KnnSearchQuery
}

func (s *KnnSearchQuery) KnnSearchQueryCaster() *KnnSearchQuery {
	_ = "STUB: not implemented"
	return nil
}
