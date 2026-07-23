package types

type KnnQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	Filter []Query `json:"filter,omitempty"`

	K *int `json:"k,omitempty"`

	NumCandidates *int    `json:"num_candidates,omitempty"`
	QueryName_    *string `json:"_name,omitempty"`

	QueryVector []float32 `json:"query_vector,omitempty"`

	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`

	RescoreVector *RescoreVector `json:"rescore_vector,omitempty"`

	Similarity *float32 `json:"similarity,omitempty"`

	VisitPercentage *float32 `json:"visit_percentage,omitempty"`
}

func (s *KnnQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewKnnQuery() *KnnQuery { _ = "STUB: not implemented"; return nil }

type KnnQueryVariant interface {
	KnnQueryCaster() *KnnQuery
}

func (s *KnnQuery) KnnQueryCaster() *KnnQuery { _ = "STUB: not implemented"; return nil }
