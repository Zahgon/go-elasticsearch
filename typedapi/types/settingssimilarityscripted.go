package types

type SettingsSimilarityScripted struct {
	Script       Script  `json:"script"`
	Type         string  `json:"type,omitempty"`
	WeightScript *Script `json:"weight_script,omitempty"`
}

func (s SettingsSimilarityScripted) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityScripted() *SettingsSimilarityScripted {
	_ = "STUB: not implemented"
	return nil
}

type SettingsSimilarityScriptedVariant interface {
	SettingsSimilarityScriptedCaster() *SettingsSimilarityScripted
}

func (s *SettingsSimilarityScripted) SettingsSimilarityScriptedCaster() *SettingsSimilarityScripted {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityScripted) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
