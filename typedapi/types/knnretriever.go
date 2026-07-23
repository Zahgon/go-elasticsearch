package types

type KnnRetriever struct {
	Field string `json:"field"`

	Filter []Query `json:"filter,omitempty"`

	K int `json:"k"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_ *string `json:"_name,omitempty"`

	NumCandidates int `json:"num_candidates"`

	QueryVector []float32 `json:"query_vector,omitempty"`

	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`

	RescoreVector *RescoreVector `json:"rescore_vector,omitempty"`

	Similarity *float32 `json:"similarity,omitempty"`

	VisitPercentage *float32 `json:"visit_percentage,omitempty"`
}

func (s *KnnRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewKnnRetriever() *KnnRetriever { _ = "STUB: not implemented"; return nil }

type KnnRetrieverVariant interface {
	KnnRetrieverCaster() *KnnRetriever
}

func (s *KnnRetriever) KnnRetrieverCaster() *KnnRetriever { _ = "STUB: not implemented"; return nil }
