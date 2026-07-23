package types

type ConnectorScheduling struct {
	Enabled bool `json:"enabled"`

	Interval string `json:"interval"`
}

func (s *ConnectorScheduling) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConnectorScheduling() *ConnectorScheduling { _ = "STUB: not implemented"; return nil }

type ConnectorSchedulingVariant interface {
	ConnectorSchedulingCaster() *ConnectorScheduling
}

func (s *ConnectorScheduling) ConnectorSchedulingCaster() *ConnectorScheduling {
	_ = "STUB: not implemented"
	return nil
}
