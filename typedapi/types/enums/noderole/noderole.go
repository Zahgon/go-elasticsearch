package noderole

type NodeRole struct {
	Name string
}

var (
	Master = NodeRole{"master"}

	Data = NodeRole{"data"}

	Datacold = NodeRole{"data_cold"}

	Datacontent = NodeRole{"data_content"}

	Datafrozen = NodeRole{"data_frozen"}

	Datahot = NodeRole{"data_hot"}

	Datawarm = NodeRole{"data_warm"}

	Client = NodeRole{"client"}

	Ingest = NodeRole{"ingest"}

	Ml = NodeRole{"ml"}

	Votingonly = NodeRole{"voting_only"}

	Transform = NodeRole{"transform"}

	Remoteclusterclient = NodeRole{"remote_cluster_client"}

	Coordinatingonly = NodeRole{"coordinating_only"}
)

func (n NodeRole) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NodeRole) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NodeRole) String() string { _ = "STUB: not implemented"; return "" }
