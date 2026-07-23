package types

type MovingAverageAggregation any

type MovingAverageAggregationVariant interface {
	MovingAverageAggregationCaster() *MovingAverageAggregation
}
