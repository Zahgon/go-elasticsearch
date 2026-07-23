package types

type AdjacencyMatrixAggregate struct {
	Buckets BucketsAdjacencyMatrixBucket `json:"buckets"`
	Meta    Metadata                     `json:"meta,omitempty"`
}

func (s *AdjacencyMatrixAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAdjacencyMatrixAggregate() *AdjacencyMatrixAggregate { _ = "STUB: not implemented"; return nil }
