package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsQuery struct {
	v *types.TermsQuery
}

func NewTermsQuery() *_termsQuery { _ = "STUB: not implemented"; return nil }

func (s *_termsQuery) TermsQuery(termsquery map[string]types.TermsQueryField) *_termsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsQuery) AddTermsQuery(key string, value types.TermsQueryFieldVariant) *_termsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsQuery) Boost(boost float32) *_termsQuery { _ = "STUB: not implemented"; return nil }

func (s *_termsQuery) QueryName_(queryname_ string) *_termsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_termsQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsQuery) TermsQueryCaster() *types.TermsQuery { _ = "STUB: not implemented"; return nil }
