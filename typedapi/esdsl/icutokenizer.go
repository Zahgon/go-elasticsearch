package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _icuTokenizer struct {
	v *types.IcuTokenizer
}

func NewIcuTokenizer(rulefiles string) *_icuTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_icuTokenizer) RuleFiles(rulefiles string) *_icuTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuTokenizer) Version(versionstring string) *_icuTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuTokenizer) IcuTokenizerCaster() *types.IcuTokenizer {
	_ = "STUB: not implemented"
	return nil
}
