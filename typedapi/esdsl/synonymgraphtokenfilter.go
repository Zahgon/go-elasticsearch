package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/synonymformat"
)

type _synonymGraphTokenFilter struct {
	v *types.SynonymGraphTokenFilter
}

func NewSynonymGraphTokenFilter() *_synonymGraphTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_synonymGraphTokenFilter) Expand(expand bool) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) Format(format synonymformat.SynonymFormat) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) Lenient(lenient bool) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) Synonyms(synonyms ...string) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) SynonymsPath(synonymspath string) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) SynonymsSet(synonymssets ...string) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) Tokenizer(tokenizer string) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) Updateable(updateable bool) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) Version(versionstring string) *_synonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymGraphTokenFilter) SynonymGraphTokenFilterCaster() *types.SynonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
