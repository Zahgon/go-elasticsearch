package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _shingleTokenFilter struct {
	v *types.ShingleTokenFilter
}

func NewShingleTokenFilter() *_shingleTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_shingleTokenFilter) FillerToken(fillertoken string) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) MaxShingleSize(stringifiedinteger types.StringifiedintegerVariant) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) MinShingleSize(stringifiedinteger types.StringifiedintegerVariant) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) OutputUnigrams(outputunigrams bool) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) OutputUnigramsIfNoShingles(outputunigramsifnoshingles bool) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) TokenSeparator(tokenseparator string) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) Version(versionstring string) *_shingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shingleTokenFilter) ShingleTokenFilterCaster() *types.ShingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
