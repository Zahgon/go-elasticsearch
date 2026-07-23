package types

type SettingsHighlight struct {
	MaxAnalyzedOffset *int `json:"max_analyzed_offset,omitempty"`
}

func (s *SettingsHighlight) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSettingsHighlight() *SettingsHighlight { _ = "STUB: not implemented"; return nil }

type SettingsHighlightVariant interface {
	SettingsHighlightCaster() *SettingsHighlight
}

func (s *SettingsHighlight) SettingsHighlightCaster() *SettingsHighlight {
	_ = "STUB: not implemented"
	return nil
}
