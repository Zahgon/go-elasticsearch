package types

type NodeInfoXpackLicenseType struct {
	Type string `json:"type"`
}

func (s *NodeInfoXpackLicenseType) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoXpackLicenseType() *NodeInfoXpackLicenseType { _ = "STUB: not implemented"; return nil }
