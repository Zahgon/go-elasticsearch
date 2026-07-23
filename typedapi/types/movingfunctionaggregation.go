package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type MovingFunctionAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Script *string `json:"script,omitempty"`

	Shift *int `json:"shift,omitempty"`

	Window *int `json:"window,omitempty"`
}

func (s *MovingFunctionAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMovingFunctionAggregation() *MovingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

type MovingFunctionAggregationVariant interface {
	MovingFunctionAggregationCaster() *MovingFunctionAggregation
}

func (s *MovingFunctionAggregation) MovingFunctionAggregationCaster() *MovingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}
