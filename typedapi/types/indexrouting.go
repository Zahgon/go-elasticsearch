package types

type IndexRouting struct {
	Allocation *IndexRoutingAllocation `json:"allocation,omitempty"`
	Rebalance  *IndexRoutingRebalance  `json:"rebalance,omitempty"`
}

func NewIndexRouting() *IndexRouting { _ = "STUB: not implemented"; return nil }

type IndexRoutingVariant interface {
	IndexRoutingCaster() *IndexRouting
}

func (s *IndexRouting) IndexRoutingCaster() *IndexRouting { _ = "STUB: not implemented"; return nil }
