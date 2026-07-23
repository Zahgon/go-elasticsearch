package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _connectorConfiguration struct {
	v types.ConnectorConfiguration
}

func NewConnectorConfiguration(connectorconfiguration map[string]types.ConnectorConfigProperties) *_connectorConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (u *_connectorConfiguration) ConnectorConfigurationCaster() *types.ConnectorConfiguration {
	_ = "STUB: not implemented"
	return nil
}
