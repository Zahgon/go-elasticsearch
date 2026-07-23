package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type MovingPercentilesAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Keyed     *bool                `json:"keyed,omitempty"`

	Shift *int `json:"shift,omitempty"`

	Window *int `json:"window,omitempty"`
}

func (s *MovingPercentilesAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMovingPercentilesAggregation() *MovingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

type MovingPercentilesAggregationVariant interface {
	MovingPercentilesAggregationCaster() *MovingPercentilesAggregation
}

func (s *MovingPercentilesAggregation) MovingPercentilesAggregationCaster() *MovingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}
