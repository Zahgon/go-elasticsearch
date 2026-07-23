package types

type SyncRulesFeature struct {
	Advanced *FeatureEnabled `json:"advanced,omitempty"`

	Basic *FeatureEnabled `json:"basic,omitempty"`
}

func NewSyncRulesFeature() *SyncRulesFeature { _ = "STUB: not implemented"; return nil }

type SyncRulesFeatureVariant interface {
	SyncRulesFeatureCaster() *SyncRulesFeature
}

func (s *SyncRulesFeature) SyncRulesFeatureCaster() *SyncRulesFeature {
	_ = "STUB: not implemented"
	return nil
}
