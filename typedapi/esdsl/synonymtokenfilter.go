package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/synonymformat"
)

type _synonymTokenFilter struct {
	v *types.SynonymTokenFilter
}

func NewSynonymTokenFilter() *_synonymTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_synonymTokenFilter) Expand(expand bool) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) Format(format synonymformat.SynonymFormat) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) Lenient(lenient bool) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) Synonyms(synonyms ...string) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) SynonymsPath(synonymspath string) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) SynonymsSet(synonymssets ...string) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) Tokenizer(tokenizer string) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) Updateable(updateable bool) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) Version(versionstring string) *_synonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymTokenFilter) SynonymTokenFilterCaster() *types.SynonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
