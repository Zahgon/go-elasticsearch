package types

type MasterIsStableIndicatorClusterFormationNode struct {
	ClusterFormationMessage string  `json:"cluster_formation_message"`
	Name                    *string `json:"name,omitempty"`
	NodeId                  string  `json:"node_id"`
}

func (s *MasterIsStableIndicatorClusterFormationNode) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMasterIsStableIndicatorClusterFormationNode() *MasterIsStableIndicatorClusterFormationNode {
	_ = "STUB: not implemented"
	return nil
}
