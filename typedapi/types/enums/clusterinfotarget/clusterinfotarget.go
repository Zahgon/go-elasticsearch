package clusterinfotarget

type ClusterInfoTarget struct {
	Name string
}

var (
	All = ClusterInfoTarget{"_all"}

	Http = ClusterInfoTarget{"http"}

	Ingest = ClusterInfoTarget{"ingest"}

	Threadpool = ClusterInfoTarget{"thread_pool"}

	Script = ClusterInfoTarget{"script"}
)

func (c ClusterInfoTarget) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterInfoTarget) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ClusterInfoTarget) String() string { _ = "STUB: not implemented"; return "" }
