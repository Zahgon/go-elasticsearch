package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticencoder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticlanguage"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticnametype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticruletype"
)

type _phoneticTokenFilter struct {
	v *types.PhoneticTokenFilter
}

func NewPhoneticTokenFilter(encoder phoneticencoder.PhoneticEncoder) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) Encoder(encoder phoneticencoder.PhoneticEncoder) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) Languageset(languagesets ...phoneticlanguage.PhoneticLanguage) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) MaxCodeLen(maxcodelen int) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) NameType(nametype phoneticnametype.PhoneticNameType) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) Replace(replace bool) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) RuleType(ruletype phoneticruletype.PhoneticRuleType) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) Version(versionstring string) *_phoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phoneticTokenFilter) PhoneticTokenFilterCaster() *types.PhoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
