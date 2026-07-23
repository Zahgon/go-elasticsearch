package types

type NodeInfoSettingsTransport struct {
	Features *NodeInfoSettingsTransportFeatures `json:"features,omitempty"`

	IgnoreDeserializationErrors Stringifiedboolean            `json:"ignore_deserialization_errors,omitempty"`
	Type                        NodeInfoSettingsTransportType `json:"type"`
	TypeDefault                 *string                       `json:"type.default,omitempty"`
}

func (s *NodeInfoSettingsTransport) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsTransport() *NodeInfoSettingsTransport {
	_ = "STUB: not implemented"
	return nil
}
