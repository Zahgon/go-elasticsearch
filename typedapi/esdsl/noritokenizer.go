package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noridecompoundmode"
)

type _noriTokenizer struct {
	v *types.NoriTokenizer
}

func NewNoriTokenizer() *_noriTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_noriTokenizer) DecompoundMode(decompoundmode noridecompoundmode.NoriDecompoundMode) *_noriTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriTokenizer) DiscardPunctuation(discardpunctuation bool) *_noriTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriTokenizer) UserDictionary(userdictionary string) *_noriTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriTokenizer) UserDictionaryRules(userdictionaryrules ...string) *_noriTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriTokenizer) Version(versionstring string) *_noriTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriTokenizer) NoriTokenizerCaster() *types.NoriTokenizer {
	_ = "STUB: not implemented"
	return nil
}
