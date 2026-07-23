package types

type IndexRoutingAllocationInclude struct {
	Id_             *string `json:"_id,omitempty"`
	TierPreference_ *string `json:"_tier_preference,omitempty"`
}

func (s *IndexRoutingAllocationInclude) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexRoutingAllocationInclude() *IndexRoutingAllocationInclude {
	_ = "STUB: not implemented"
	return nil
}

type IndexRoutingAllocationIncludeVariant interface {
	IndexRoutingAllocationIncludeCaster() *IndexRoutingAllocationInclude
}

func (s *IndexRoutingAllocationInclude) IndexRoutingAllocationIncludeCaster() *IndexRoutingAllocationInclude {
	_ = "STUB: not implemented"
	return nil
}
