package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _connectorFeatures struct {
	v *types.ConnectorFeatures
}

func NewConnectorFeatures() *_connectorFeatures { _ = "STUB: not implemented"; return nil }

func (s *_connectorFeatures) DocumentLevelSecurity(documentlevelsecurity types.FeatureEnabledVariant) *_connectorFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorFeatures) IncrementalSync(incrementalsync types.FeatureEnabledVariant) *_connectorFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorFeatures) NativeConnectorApiKeys(nativeconnectorapikeys types.FeatureEnabledVariant) *_connectorFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorFeatures) SyncRules(syncrules types.SyncRulesFeatureVariant) *_connectorFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorFeatures) ConnectorFeaturesCaster() *types.ConnectorFeatures {
	_ = "STUB: not implemented"
	return nil
}
