package types

type IndexRoutingAllocationDisk struct {
	ThresholdEnabled *string `json:"threshold_enabled,omitempty"`
}

func (s *IndexRoutingAllocationDisk) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexRoutingAllocationDisk() *IndexRoutingAllocationDisk {
	_ = "STUB: not implemented"
	return nil
}

type IndexRoutingAllocationDiskVariant interface {
	IndexRoutingAllocationDiskCaster() *IndexRoutingAllocationDisk
}

func (s *IndexRoutingAllocationDisk) IndexRoutingAllocationDiskCaster() *IndexRoutingAllocationDisk {
	_ = "STUB: not implemented"
	return nil
}
