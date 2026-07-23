package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/kuromojitokenizationmode"
)

type _kuromojiTokenizer struct {
	v *types.KuromojiTokenizer
}

func NewKuromojiTokenizer(mode kuromojitokenizationmode.KuromojiTokenizationMode) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) DiscardCompoundToken(discardcompoundtoken bool) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) DiscardPunctuation(discardpunctuation bool) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) Mode(mode kuromojitokenizationmode.KuromojiTokenizationMode) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) NbestCost(nbestcost int) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) NbestExamples(nbestexamples string) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) UserDictionary(userdictionary string) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) UserDictionaryRules(userdictionaryrules ...string) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) Version(versionstring string) *_kuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kuromojiTokenizer) KuromojiTokenizerCaster() *types.KuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}
