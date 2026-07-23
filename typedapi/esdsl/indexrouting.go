package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexRouting struct {
	v *types.IndexRouting
}

func NewIndexRouting() *_indexRouting { _ = "STUB: not implemented"; return nil }

func (s *_indexRouting) Allocation(allocation types.IndexRoutingAllocationVariant) *_indexRouting {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRouting) Rebalance(rebalance types.IndexRoutingRebalanceVariant) *_indexRouting {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRouting) IndexRoutingCaster() *types.IndexRouting {
	_ = "STUB: not implemented"
	return nil
}
