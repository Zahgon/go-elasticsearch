package types

type AdjacencyMatrixBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          string               `json:"key"`
}

func (s *AdjacencyMatrixBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s AdjacencyMatrixBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewAdjacencyMatrixBucket() *AdjacencyMatrixBucket { _ = "STUB: not implemented"; return nil }
