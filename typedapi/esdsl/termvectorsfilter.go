package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termVectorsFilter struct {
	v *types.TermVectorsFilter
}

func NewTermVectorsFilter() *_termVectorsFilter { _ = "STUB: not implemented"; return nil }

func (s *_termVectorsFilter) MaxDocFreq(maxdocfreq int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) MaxNumTerms(maxnumterms int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) MaxTermFreq(maxtermfreq int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) MaxWordLength(maxwordlength int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) MinDocFreq(mindocfreq int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) MinTermFreq(mintermfreq int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) MinWordLength(minwordlength int) *_termVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termVectorsFilter) TermVectorsFilterCaster() *types.TermVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}
