package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _multiplexerTokenFilter struct {
	v *types.MultiplexerTokenFilter
}

func NewMultiplexerTokenFilter() *_multiplexerTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_multiplexerTokenFilter) Filters(filters ...string) *_multiplexerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiplexerTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_multiplexerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiplexerTokenFilter) Version(versionstring string) *_multiplexerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiplexerTokenFilter) MultiplexerTokenFilterCaster() *types.MultiplexerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
