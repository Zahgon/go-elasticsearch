package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestsort"
)

type _termSuggester struct {
	v *types.TermSuggester
}

func NewTermSuggester() *_termSuggester { _ = "STUB: not implemented"; return nil }

func (s *_termSuggester) LowercaseTerms(lowercaseterms bool) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) MaxEdits(maxedits int) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) MaxInspections(maxinspections int) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) MaxTermFreq(maxtermfreq float32) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) MinDocFreq(mindocfreq float32) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) MinWordLength(minwordlength int) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) PrefixLength(prefixlength int) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) ShardSize(shardsize int) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) Sort(sort suggestsort.SuggestSort) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) StringDistance(stringdistance stringdistance.StringDistance) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) SuggestMode(suggestmode suggestmode.SuggestMode) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) Analyzer(analyzer string) *_termSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) Field(field string) *_termSuggester { _ = "STUB: not implemented"; return nil }

func (s *_termSuggester) Size(size int) *_termSuggester { _ = "STUB: not implemented"; return nil }

func (s *_termSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termSuggester) TermSuggesterCaster() *types.TermSuggester {
	_ = "STUB: not implemented"
	return nil
}
