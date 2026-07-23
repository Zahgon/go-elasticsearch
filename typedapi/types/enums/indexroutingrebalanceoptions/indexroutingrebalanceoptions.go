package indexroutingrebalanceoptions

type IndexRoutingRebalanceOptions struct {
	Name string
}

var (
	All = IndexRoutingRebalanceOptions{"all"}

	Primaries = IndexRoutingRebalanceOptions{"primaries"}

	Replicas = IndexRoutingRebalanceOptions{"replicas"}

	None = IndexRoutingRebalanceOptions{"none"}
)

func (i IndexRoutingRebalanceOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexRoutingRebalanceOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndexRoutingRebalanceOptions) String() string { _ = "STUB: not implemented"; return "" }
