package connectorstatus

type ConnectorStatus struct {
	Name string
}

var (
	Created = ConnectorStatus{"created"}

	Needsconfiguration = ConnectorStatus{"needs_configuration"}

	Configured = ConnectorStatus{"configured"}

	Connected = ConnectorStatus{"connected"}

	Error = ConnectorStatus{"error"}
)

func (c ConnectorStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnectorStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ConnectorStatus) String() string { _ = "STUB: not implemented"; return "" }
