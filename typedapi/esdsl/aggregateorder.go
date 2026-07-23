package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _aggregateOrder struct {
	v types.AggregateOrder
}

func NewAggregateOrder() *_aggregateOrder { _ = "STUB: not implemented"; return nil }

func (u *_aggregateOrder) Map(value map[string]sortorder.SortOrder) *_aggregateOrder {
	_ = "STUB: not implemented"
	return nil
}

func (u *_aggregateOrder) SortOrders(sortorders ...map[string]sortorder.SortOrder) *_aggregateOrder {
	_ = "STUB: not implemented"
	return nil
}

func (u *_aggregateOrder) AggregateOrderCaster() *types.AggregateOrder {
	_ = "STUB: not implemented"
	return nil
}
