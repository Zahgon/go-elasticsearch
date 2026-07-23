package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _phraseSuggestHighlight struct {
	v *types.PhraseSuggestHighlight
}

func NewPhraseSuggestHighlight(posttag string, pretag string) *_phraseSuggestHighlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestHighlight) PostTag(posttag string) *_phraseSuggestHighlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestHighlight) PreTag(pretag string) *_phraseSuggestHighlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestHighlight) PhraseSuggestHighlightCaster() *types.PhraseSuggestHighlight {
	_ = "STUB: not implemented"
	return nil
}
