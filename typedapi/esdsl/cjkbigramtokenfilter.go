package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cjkbigramignoredscript"
)

type _cjkBigramTokenFilter struct {
	v *types.CjkBigramTokenFilter
}

func NewCjkBigramTokenFilter() *_cjkBigramTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_cjkBigramTokenFilter) IgnoredScripts(ignoredscripts ...cjkbigramignoredscript.CjkBigramIgnoredScript) *_cjkBigramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cjkBigramTokenFilter) OutputUnigrams(outputunigrams bool) *_cjkBigramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cjkBigramTokenFilter) Version(versionstring string) *_cjkBigramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cjkBigramTokenFilter) CjkBigramTokenFilterCaster() *types.CjkBigramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
