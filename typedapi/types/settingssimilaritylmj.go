package types

type SettingsSimilarityLmj struct {
	Lambda *Float64 `json:"lambda,omitempty"`
	Type   string   `json:"type,omitempty"`
}

func (s *SettingsSimilarityLmj) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SettingsSimilarityLmj) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityLmj() *SettingsSimilarityLmj { _ = "STUB: not implemented"; return nil }

type SettingsSimilarityLmjVariant interface {
	SettingsSimilarityLmjCaster() *SettingsSimilarityLmj
}

func (s *SettingsSimilarityLmj) SettingsSimilarityLmjCaster() *SettingsSimilarityLmj {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityLmj) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
