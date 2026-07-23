package indexroutingallocationoptions

type IndexRoutingAllocationOptions struct {
	Name string
}

var (
	All = IndexRoutingAllocationOptions{"all"}

	Primaries = IndexRoutingAllocationOptions{"primaries"}

	Newprimaries = IndexRoutingAllocationOptions{"new_primaries"}

	None = IndexRoutingAllocationOptions{"none"}
)

func (i IndexRoutingAllocationOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexRoutingAllocationOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndexRoutingAllocationOptions) String() string { _ = "STUB: not implemented"; return "" }
