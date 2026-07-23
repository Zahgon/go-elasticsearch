package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/kuromojitokenizationmode"
)

type _kuromojiAnalyzer struct {
	v *types.KuromojiAnalyzer
}

func NewKuromojiAnalyzer() *_kuromojiAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_kuromojiAnalyzer) Mode(mode kuromojitokenizationmode.KuromojiTokenizationMode) *_kuromojiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiAnalyzer) UserDictionary(userdictionary string) *_kuromojiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiAnalyzer) KuromojiAnalyzerCaster() *types.KuromojiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
