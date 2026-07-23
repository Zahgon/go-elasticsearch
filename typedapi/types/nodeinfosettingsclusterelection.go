package types

type NodeInfoSettingsClusterElection struct {
	Strategy string `json:"strategy"`
}

func (s *NodeInfoSettingsClusterElection) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsClusterElection() *NodeInfoSettingsClusterElection {
	_ = "STUB: not implemented"
	return nil
}
