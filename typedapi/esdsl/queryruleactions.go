package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _queryRuleActions struct {
	v *types.QueryRuleActions
}

func NewQueryRuleActions() *_queryRuleActions { _ = "STUB: not implemented"; return nil }

func (s *_queryRuleActions) Docs(docs ...types.PinnedDocVariant) *_queryRuleActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleActions) DocsValues(docsvalues []types.PinnedDoc) *_queryRuleActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleActions) Ids(ids ...string) *_queryRuleActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleActions) QueryRuleActionsCaster() *types.QueryRuleActions {
	_ = "STUB: not implemented"
	return nil
}
