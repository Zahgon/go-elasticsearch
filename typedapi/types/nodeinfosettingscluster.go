package types

type NodeInfoSettingsCluster struct {
	DeprecationIndexing *DeprecationIndexing            `json:"deprecation_indexing,omitempty"`
	Election            NodeInfoSettingsClusterElection `json:"election"`
	InitialMasterNodes  []string                        `json:"initial_master_nodes,omitempty"`
	Name                string                          `json:"name"`
	Routing             *IndexRouting                   `json:"routing,omitempty"`
}

func (s *NodeInfoSettingsCluster) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsCluster() *NodeInfoSettingsCluster { _ = "STUB: not implemented"; return nil }
