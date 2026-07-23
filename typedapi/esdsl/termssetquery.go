package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsSetQuery struct {
	k string
	v *types.TermsSetQuery
}

func NewTermsSetQuery(key string) *_termsSetQuery { _ = "STUB: not implemented"; return nil }

func (s *_termsSetQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) MinimumShouldMatchField(field string) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) MinimumShouldMatchScript(minimumshouldmatchscript types.ScriptVariant) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) Terms(terms ...types.FieldValueVariant) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) TermsValues(termsvalues []types.FieldValue) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) Boost(boost float32) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) QueryName_(queryname_ string) *_termsSetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsSetQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleTermsSetQuery() *_termsSetQuery { _ = "STUB: not implemented"; return nil }

func (s *_termsSetQuery) TermsSetQueryCaster() *types.TermsSetQuery {
	_ = "STUB: not implemented"
	return nil
}
