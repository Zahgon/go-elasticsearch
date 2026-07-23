package nodesinfometric

type NodesInfoMetric struct {
	Name string
}

var (
	All = NodesInfoMetric{"_all"}

	None = NodesInfoMetric{"_none"}

	Settings = NodesInfoMetric{"settings"}

	Os = NodesInfoMetric{"os"}

	Process = NodesInfoMetric{"process"}

	Jvm = NodesInfoMetric{"jvm"}

	Threadpool = NodesInfoMetric{"thread_pool"}

	Transport = NodesInfoMetric{"transport"}

	Http = NodesInfoMetric{"http"}

	Remoteclusterserver = NodesInfoMetric{"remote_cluster_server"}

	Plugins = NodesInfoMetric{"plugins"}

	Ingest = NodesInfoMetric{"ingest"}

	Aggregations = NodesInfoMetric{"aggregations"}

	Indices = NodesInfoMetric{"indices"}
)

func (n NodesInfoMetric) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NodesInfoMetric) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NodesInfoMetric) String() string { _ = "STUB: not implemented"; return "" }
