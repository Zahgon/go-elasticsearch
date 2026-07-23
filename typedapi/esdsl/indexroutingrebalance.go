package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexroutingrebalanceoptions"
)

type _indexRoutingRebalance struct {
	v *types.IndexRoutingRebalance
}

func NewIndexRoutingRebalance(enable indexroutingrebalanceoptions.IndexRoutingRebalanceOptions) *_indexRoutingRebalance {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRoutingRebalance) Enable(enable indexroutingrebalanceoptions.IndexRoutingRebalanceOptions) *_indexRoutingRebalance {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRoutingRebalance) IndexRoutingRebalanceCaster() *types.IndexRoutingRebalance {
	_ = "STUB: not implemented"
	return nil
}
