package types

type SettingsSimilarityBoolean struct {
	Type string `json:"type,omitempty"`
}

func (s SettingsSimilarityBoolean) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityBoolean() *SettingsSimilarityBoolean {
	_ = "STUB: not implemented"
	return nil
}

type SettingsSimilarityBooleanVariant interface {
	SettingsSimilarityBooleanCaster() *SettingsSimilarityBoolean
}

func (s *SettingsSimilarityBoolean) SettingsSimilarityBooleanCaster() *SettingsSimilarityBoolean {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityBoolean) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
