package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/edgengramside"
)

type _edgeNGramTokenFilter struct {
	v *types.EdgeNGramTokenFilter
}

func NewEdgeNGramTokenFilter() *_edgeNGramTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_edgeNGramTokenFilter) MaxGram(maxgram int) *_edgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenFilter) MinGram(mingram int) *_edgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_edgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenFilter) Side(side edgengramside.EdgeNGramSide) *_edgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenFilter) Version(versionstring string) *_edgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenFilter) EdgeNGramTokenFilterCaster() *types.EdgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
