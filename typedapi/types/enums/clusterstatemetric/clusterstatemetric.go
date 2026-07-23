package clusterstatemetric

type ClusterStateMetric struct {
	Name string
}

var (
	All = ClusterStateMetric{"_all"}

	Version = ClusterStateMetric{"version"}

	Masternode = ClusterStateMetric{"master_node"}

	Blocks = ClusterStateMetric{"blocks"}

	Nodes = ClusterStateMetric{"nodes"}

	Metadata = ClusterStateMetric{"metadata"}

	Routingtable = ClusterStateMetric{"routing_table"}

	Routingnodes = ClusterStateMetric{"routing_nodes"}

	Customs = ClusterStateMetric{"customs"}
)

func (c ClusterStateMetric) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterStateMetric) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c ClusterStateMetric) String() string { _ = "STUB: not implemented"; return "" }
