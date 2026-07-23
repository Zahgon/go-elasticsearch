package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textExpansionQuery struct {
	k string
	v *types.TextExpansionQuery
}

func NewTextExpansionQuery(key string) *_textExpansionQuery { _ = "STUB: not implemented"; return nil }

func (s *_textExpansionQuery) ModelId(modelid string) *_textExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionQuery) ModelText(modeltext string) *_textExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_textExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionQuery) Boost(boost float32) *_textExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionQuery) QueryName_(queryname_ string) *_textExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleTextExpansionQuery() *_textExpansionQuery { _ = "STUB: not implemented"; return nil }

func (s *_textExpansionQuery) TextExpansionQueryCaster() *types.TextExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}
