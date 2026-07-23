package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termQuery struct {
	k string
	v *types.TermQuery
}

func NewTermQuery(field string, value types.FieldValueVariant) *_termQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termQuery) CaseInsensitive(caseinsensitive bool) *_termQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termQuery) Value(fieldvalue types.FieldValueVariant) *_termQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termQuery) Boost(boost float32) *_termQuery { _ = "STUB: not implemented"; return nil }

func (s *_termQuery) QueryName_(queryname_ string) *_termQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_termQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleTermQuery() *_termQuery { _ = "STUB: not implemented"; return nil }

func (s *_termQuery) TermQueryCaster() *types.TermQuery { _ = "STUB: not implemented"; return nil }
