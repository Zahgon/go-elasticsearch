package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/kuromojitokenizationmode"
)

type KuromojiAnalyzer struct {
	Mode           *kuromojitokenizationmode.KuromojiTokenizationMode `json:"mode,omitempty"`
	Type           string                                             `json:"type,omitempty"`
	UserDictionary *string                                            `json:"user_dictionary,omitempty"`
}

func (s *KuromojiAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s KuromojiAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewKuromojiAnalyzer() *KuromojiAnalyzer { _ = "STUB: not implemented"; return nil }

type KuromojiAnalyzerVariant interface {
	KuromojiAnalyzerCaster() *KuromojiAnalyzer
}

func (s *KuromojiAnalyzer) KuromojiAnalyzerCaster() *KuromojiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *KuromojiAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
