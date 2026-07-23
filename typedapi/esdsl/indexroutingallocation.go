package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexroutingallocationoptions"
)

type _indexRoutingAllocation struct {
	v *types.IndexRoutingAllocation
}

func NewIndexRoutingAllocation() *_indexRoutingAllocation { _ = "STUB: not implemented"; return nil }

func (s *_indexRoutingAllocation) Disk(disk types.IndexRoutingAllocationDiskVariant) *_indexRoutingAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRoutingAllocation) Enable(enable indexroutingallocationoptions.IndexRoutingAllocationOptions) *_indexRoutingAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRoutingAllocation) Include(include types.IndexRoutingAllocationIncludeVariant) *_indexRoutingAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRoutingAllocation) InitialRecovery(initialrecovery types.IndexRoutingAllocationInitialRecoveryVariant) *_indexRoutingAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexRoutingAllocation) IndexRoutingAllocationCaster() *types.IndexRoutingAllocation {
	_ = "STUB: not implemented"
	return nil
}
