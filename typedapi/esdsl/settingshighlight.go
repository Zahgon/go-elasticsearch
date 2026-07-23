package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settingsHighlight struct {
	v *types.SettingsHighlight
}

func NewSettingsHighlight() *_settingsHighlight { _ = "STUB: not implemented"; return nil }

func (s *_settingsHighlight) MaxAnalyzedOffset(maxanalyzedoffset int) *_settingsHighlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsHighlight) SettingsHighlightCaster() *types.SettingsHighlight {
	_ = "STUB: not implemented"
	return nil
}
