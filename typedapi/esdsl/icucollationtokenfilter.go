package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationstrength"
)

type _icuCollationTokenFilter struct {
	v *types.IcuCollationTokenFilter
}

func NewIcuCollationTokenFilter() *_icuCollationTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_icuCollationTokenFilter) Alternate(alternate icucollationalternate.IcuCollationAlternate) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) CaseFirst(casefirst icucollationcasefirst.IcuCollationCaseFirst) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) CaseLevel(caselevel bool) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Country(country string) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Decomposition(decomposition icucollationdecomposition.IcuCollationDecomposition) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) HiraganaQuaternaryMode(hiraganaquaternarymode bool) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Language(language string) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Numeric(numeric bool) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Rules(rules string) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Strength(strength icucollationstrength.IcuCollationStrength) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) VariableTop(variabletop string) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Variant(variant string) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) Version(versionstring string) *_icuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationTokenFilter) IcuCollationTokenFilterCaster() *types.IcuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
