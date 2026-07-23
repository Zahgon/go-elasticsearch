package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _moreLikeThisQuery struct {
	v *types.MoreLikeThisQuery
}

func NewMoreLikeThisQuery() *_moreLikeThisQuery { _ = "STUB: not implemented"; return nil }

func (s *_moreLikeThisQuery) Analyzer(analyzer string) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) BoostTerms(boostterms types.Float64) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) FailOnUnsupportedField(failonunsupportedfield bool) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Fields(fields ...string) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Include(include bool) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Like(likes ...types.LikeVariant) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MaxDocFreq(maxdocfreq int) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MaxQueryTerms(maxqueryterms int) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MaxWordLength(maxwordlength int) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MinDocFreq(mindocfreq int) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MinTermFreq(mintermfreq int) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MinWordLength(minwordlength int) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Routing(routing string) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) StopWords(stopwords types.StopWordsVariant) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Unlike(unlikes ...types.LikeVariant) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Version(versionnumber int64) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) VersionType(versiontype versiontype.VersionType) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) Boost(boost float32) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) QueryName_(queryname_ string) *_moreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_moreLikeThisQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_moreLikeThisQuery) MoreLikeThisQueryCaster() *types.MoreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}
