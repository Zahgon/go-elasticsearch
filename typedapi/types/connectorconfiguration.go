package types

type ConnectorConfiguration map[string]ConnectorConfigProperties

type ConnectorConfigurationVariant interface {
	ConnectorConfigurationCaster() *ConnectorConfiguration
}
