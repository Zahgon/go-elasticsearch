package types

type AggregateOrder any

type AggregateOrderVariant interface {
	AggregateOrderCaster() *AggregateOrder
}
