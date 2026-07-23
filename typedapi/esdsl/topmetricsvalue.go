package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _topMetricsValue struct {
	v *types.TopMetricsValue
}

func NewTopMetricsValue() *_topMetricsValue { _ = "STUB: not implemented"; return nil }

func (s *_topMetricsValue) Field(field string) *_topMetricsValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsValue) TopMetricsValueCaster() *types.TopMetricsValue {
	_ = "STUB: not implemented"
	return nil
}
