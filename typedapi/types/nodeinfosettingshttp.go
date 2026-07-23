package types

type NodeInfoSettingsHttp struct {
	Compression *string                  `json:"compression,omitempty"`
	Port        *string                  `json:"port,omitempty"`
	Type        NodeInfoSettingsHttpType `json:"type"`
	TypeDefault *string                  `json:"type.default,omitempty"`
}

func (s *NodeInfoSettingsHttp) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsHttp() *NodeInfoSettingsHttp { _ = "STUB: not implemented"; return nil }
