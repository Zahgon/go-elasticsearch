package types

type SettingsSimilarityBm25 struct {
	B                *Float64 `json:"b,omitempty"`
	DiscountOverlaps *bool    `json:"discount_overlaps,omitempty"`
	K1               *Float64 `json:"k1,omitempty"`
	Type             string   `json:"type,omitempty"`
}

func (s *SettingsSimilarityBm25) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SettingsSimilarityBm25) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityBm25() *SettingsSimilarityBm25 { _ = "STUB: not implemented"; return nil }

type SettingsSimilarityBm25Variant interface {
	SettingsSimilarityBm25Caster() *SettingsSimilarityBm25
}

func (s *SettingsSimilarityBm25) SettingsSimilarityBm25Caster() *SettingsSimilarityBm25 {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityBm25) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
