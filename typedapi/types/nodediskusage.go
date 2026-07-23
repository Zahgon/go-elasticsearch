package types

type NodeDiskUsage struct {
	LeastAvailable DiskUsage `json:"least_available"`
	MostAvailable  DiskUsage `json:"most_available"`
	NodeName       string    `json:"node_name"`
}

func (s *NodeDiskUsage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeDiskUsage() *NodeDiskUsage { _ = "STUB: not implemented"; return nil }
