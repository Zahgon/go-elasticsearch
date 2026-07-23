package nodesusagemetric

type NodesUsageMetric struct {
	Name string
}

var (
	All = NodesUsageMetric{"_all"}

	Restactions = NodesUsageMetric{"rest_actions"}

	Aggregations = NodesUsageMetric{"aggregations"}
)

func (n NodesUsageMetric) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NodesUsageMetric) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NodesUsageMetric) String() string { _ = "STUB: not implemented"; return "" }
