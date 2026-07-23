package types

type NodeInfoSettingsTransportFeatures struct {
	XPack string `json:"x-pack"`
}

func (s *NodeInfoSettingsTransportFeatures) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsTransportFeatures() *NodeInfoSettingsTransportFeatures {
	_ = "STUB: not implemented"
	return nil
}
