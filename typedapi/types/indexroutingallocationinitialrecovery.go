package types

type IndexRoutingAllocationInitialRecovery struct {
	Id_ *string `json:"_id,omitempty"`
}

func (s *IndexRoutingAllocationInitialRecovery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexRoutingAllocationInitialRecovery() *IndexRoutingAllocationInitialRecovery {
	_ = "STUB: not implemented"
	return nil
}

type IndexRoutingAllocationInitialRecoveryVariant interface {
	IndexRoutingAllocationInitialRecoveryCaster() *IndexRoutingAllocationInitialRecovery
}

func (s *IndexRoutingAllocationInitialRecovery) IndexRoutingAllocationInitialRecoveryCaster() *IndexRoutingAllocationInitialRecovery {
	_ = "STUB: not implemented"
	return nil
}
