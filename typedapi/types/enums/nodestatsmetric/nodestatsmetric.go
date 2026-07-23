package nodestatsmetric

type NodeStatsMetric struct {
	Name string
}

var (
	All = NodeStatsMetric{"_all"}

	None = NodeStatsMetric{"_none"}

	Indices = NodeStatsMetric{"indices"}

	Os = NodeStatsMetric{"os"}

	Process = NodeStatsMetric{"process"}

	Jvm = NodeStatsMetric{"jvm"}

	Threadpool = NodeStatsMetric{"thread_pool"}

	Fs = NodeStatsMetric{"fs"}

	Transport = NodeStatsMetric{"transport"}

	Http = NodeStatsMetric{"http"}

	Breaker = NodeStatsMetric{"breaker"}

	Script = NodeStatsMetric{"script"}

	Discovery = NodeStatsMetric{"discovery"}

	Ingest = NodeStatsMetric{"ingest"}

	Adaptiveselection = NodeStatsMetric{"adaptive_selection"}

	Scriptcache = NodeStatsMetric{"script_cache"}

	Indexingpressure = NodeStatsMetric{"indexing_pressure"}

	Repositories = NodeStatsMetric{"repositories"}

	Allocations = NodeStatsMetric{"allocations"}
)

func (n NodeStatsMetric) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NodeStatsMetric) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NodeStatsMetric) String() string { _ = "STUB: not implemented"; return "" }
