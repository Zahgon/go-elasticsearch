package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
)

type _directGenerator struct {
	v *types.DirectGenerator
}

func NewDirectGenerator() *_directGenerator { _ = "STUB: not implemented"; return nil }

func (s *_directGenerator) Field(field string) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) MaxEdits(maxedits int) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) MaxInspections(maxinspections float32) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) MaxTermFreq(maxtermfreq float32) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) MinDocFreq(mindocfreq float32) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) MinWordLength(minwordlength int) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) PostFilter(postfilter string) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) PreFilter(prefilter string) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) PrefixLength(prefixlength int) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) Size(size int) *_directGenerator { _ = "STUB: not implemented"; return nil }

func (s *_directGenerator) SuggestMode(suggestmode suggestmode.SuggestMode) *_directGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *_directGenerator) DirectGeneratorCaster() *types.DirectGenerator {
	_ = "STUB: not implemented"
	return nil
}
