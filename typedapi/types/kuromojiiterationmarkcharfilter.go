package types

type KuromojiIterationMarkCharFilter struct {
	NormalizeKana  bool    `json:"normalize_kana"`
	NormalizeKanji bool    `json:"normalize_kanji"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *KuromojiIterationMarkCharFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KuromojiIterationMarkCharFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKuromojiIterationMarkCharFilter() *KuromojiIterationMarkCharFilter {
	_ = "STUB: not implemented"
	return nil
}

type KuromojiIterationMarkCharFilterVariant interface {
	KuromojiIterationMarkCharFilterCaster() *KuromojiIterationMarkCharFilter
}

func (s *KuromojiIterationMarkCharFilter) KuromojiIterationMarkCharFilterCaster() *KuromojiIterationMarkCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KuromojiIterationMarkCharFilter) CharFilterDefinitionCaster() *CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
