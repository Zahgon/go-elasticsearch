package types

type ConnectorFeatures struct {
	DocumentLevelSecurity *FeatureEnabled `json:"document_level_security,omitempty"`

	IncrementalSync *FeatureEnabled `json:"incremental_sync,omitempty"`

	NativeConnectorApiKeys *FeatureEnabled   `json:"native_connector_api_keys,omitempty"`
	SyncRules              *SyncRulesFeature `json:"sync_rules,omitempty"`
}

func NewConnectorFeatures() *ConnectorFeatures { _ = "STUB: not implemented"; return nil }

type ConnectorFeaturesVariant interface {
	ConnectorFeaturesCaster() *ConnectorFeatures
}

func (s *ConnectorFeatures) ConnectorFeaturesCaster() *ConnectorFeatures {
	_ = "STUB: not implemented"
	return nil
}
