package nodestatslevel

type NodeStatsLevel struct {
	Name string
}

var (
	Node = NodeStatsLevel{"node"}

	Indices = NodeStatsLevel{"indices"}

	Shards = NodeStatsLevel{"shards"}
)

func (n NodeStatsLevel) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NodeStatsLevel) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NodeStatsLevel) String() string { _ = "STUB: not implemented"; return "" }
