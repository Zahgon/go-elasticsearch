package types

type AdjacencyMatrixAggregation struct {
	Filters map[string]Query `json:"filters,omitempty"`

	Separator *string `json:"separator,omitempty"`
}

func (s *AdjacencyMatrixAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAdjacencyMatrixAggregation() *AdjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}

type AdjacencyMatrixAggregationVariant interface {
	AdjacencyMatrixAggregationCaster() *AdjacencyMatrixAggregation
}

func (s *AdjacencyMatrixAggregation) AdjacencyMatrixAggregationCaster() *AdjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}
