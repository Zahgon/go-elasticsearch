package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _tokenFilter struct {
	v types.TokenFilter
}

func NewTokenFilter() *_tokenFilter { _ = "STUB: not implemented"; return nil }

func (u *_tokenFilter) String(string string) *_tokenFilter { _ = "STUB: not implemented"; return nil }

func (u *_tokenFilter) TokenFilterDefinition(tokenfilterdefinition types.TokenFilterDefinitionVariant) *_tokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_tokenFilterDefinition) TokenFilterCaster() *types.TokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_tokenFilter) TokenFilterCaster() *types.TokenFilter {
	_ = "STUB: not implemented"
	return nil
}
