package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/lowercasetokenfilterlanguages"
)

type _lowercaseTokenFilter struct {
	v *types.LowercaseTokenFilter
}

func NewLowercaseTokenFilter() *_lowercaseTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_lowercaseTokenFilter) Language(language lowercasetokenfilterlanguages.LowercaseTokenFilterLanguages) *_lowercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseTokenFilter) Version(versionstring string) *_lowercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseTokenFilter) LowercaseTokenFilterCaster() *types.LowercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
