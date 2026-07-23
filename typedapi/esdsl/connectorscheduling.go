package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _connectorScheduling struct {
	v *types.ConnectorScheduling
}

func NewConnectorScheduling(enabled bool, interval string) *_connectorScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorScheduling) Enabled(enabled bool) *_connectorScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorScheduling) Interval(interval string) *_connectorScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorScheduling) ConnectorSchedulingCaster() *types.ConnectorScheduling {
	_ = "STUB: not implemented"
	return nil
}
