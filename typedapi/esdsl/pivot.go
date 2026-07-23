package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pivot struct {
	v *types.Pivot
}

func NewPivot() *_pivot { _ = "STUB: not implemented"; return nil }

func (s *_pivot) Aggregations(aggregations map[string]types.Aggregations) *_pivot {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivot) AddAggregation(key string, value types.AggregationsVariant) *_pivot {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivot) GroupBy(groupby map[string]types.PivotGroupByContainer) *_pivot {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivot) AddGroupBy(key string, value types.PivotGroupByContainerVariant) *_pivot {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivot) PivotCaster() *types.Pivot { _ = "STUB: not implemented"; return nil }
