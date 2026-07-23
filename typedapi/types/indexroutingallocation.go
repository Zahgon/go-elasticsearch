package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexroutingallocationoptions"
)

type IndexRoutingAllocation struct {
	Disk            *IndexRoutingAllocationDisk                                  `json:"disk,omitempty"`
	Enable          *indexroutingallocationoptions.IndexRoutingAllocationOptions `json:"enable,omitempty"`
	Include         *IndexRoutingAllocationInclude                               `json:"include,omitempty"`
	InitialRecovery *IndexRoutingAllocationInitialRecovery                       `json:"initial_recovery,omitempty"`
}

func NewIndexRoutingAllocation() *IndexRoutingAllocation { _ = "STUB: not implemented"; return nil }

type IndexRoutingAllocationVariant interface {
	IndexRoutingAllocationCaster() *IndexRoutingAllocation
}

func (s *IndexRoutingAllocation) IndexRoutingAllocationCaster() *IndexRoutingAllocation {
	_ = "STUB: not implemented"
	return nil
}
