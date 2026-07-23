package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _asciiFoldingTokenFilter struct {
	v *types.AsciiFoldingTokenFilter
}

func NewAsciiFoldingTokenFilter() *_asciiFoldingTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_asciiFoldingTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_asciiFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_asciiFoldingTokenFilter) Version(versionstring string) *_asciiFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_asciiFoldingTokenFilter) AsciiFoldingTokenFilterCaster() *types.AsciiFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
