package types

type SettingsSimilarityLmd struct {
	Mu   *Float64 `json:"mu,omitempty"`
	Type string   `json:"type,omitempty"`
}

func (s *SettingsSimilarityLmd) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SettingsSimilarityLmd) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityLmd() *SettingsSimilarityLmd { _ = "STUB: not implemented"; return nil }

type SettingsSimilarityLmdVariant interface {
	SettingsSimilarityLmdCaster() *SettingsSimilarityLmd
}

func (s *SettingsSimilarityLmd) SettingsSimilarityLmdCaster() *SettingsSimilarityLmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityLmd) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
