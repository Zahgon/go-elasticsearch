package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _adjacencyMatrixAggregation struct {
	v *types.AdjacencyMatrixAggregation
}

func NewAdjacencyMatrixAggregation() *_adjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adjacencyMatrixAggregation) Filters(filters map[string]types.Query) *_adjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adjacencyMatrixAggregation) AddFilter(key string, value types.QueryVariant) *_adjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adjacencyMatrixAggregation) Separator(separator string) *_adjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adjacencyMatrixAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adjacencyMatrixAggregation) AdjacencyMatrixAggregationCaster() *types.AdjacencyMatrixAggregation {
	_ = "STUB: not implemented"
	return nil
}
