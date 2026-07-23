package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexroutingrebalanceoptions"
)

type IndexRoutingRebalance struct {
	Enable indexroutingrebalanceoptions.IndexRoutingRebalanceOptions `json:"enable"`
}

func NewIndexRoutingRebalance() *IndexRoutingRebalance { _ = "STUB: not implemented"; return nil }

type IndexRoutingRebalanceVariant interface {
	IndexRoutingRebalanceCaster() *IndexRoutingRebalance
}

func (s *IndexRoutingRebalance) IndexRoutingRebalanceCaster() *IndexRoutingRebalance {
	_ = "STUB: not implemented"
	return nil
}
