package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/stopwordlanguage"
)

type _stopWords struct {
	v types.StopWords
}

func NewStopWords() *_stopWords { _ = "STUB: not implemented"; return nil }

func (u *_stopWords) StopWordLanguage(stopwordlanguage stopwordlanguage.StopWordLanguage) *_stopWords {
	_ = "STUB: not implemented"
	return nil
}

func (u *_stopWords) Strings(strings ...string) *_stopWords { _ = "STUB: not implemented"; return nil }

func (u *_stopWords) StopWordsCaster() *types.StopWords { _ = "STUB: not implemented"; return nil }
