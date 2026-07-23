package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptScoreQuery struct {
	v *types.ScriptScoreQuery
}

func NewScriptScoreQuery(query types.QueryVariant, script types.ScriptVariant) *_scriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreQuery) MinScore(minscore float32) *_scriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreQuery) Query(query types.QueryVariant) *_scriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreQuery) Script(script types.ScriptVariant) *_scriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreQuery) Boost(boost float32) *_scriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreQuery) QueryName_(queryname_ string) *_scriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_scriptScoreQuery) ScriptScoreQueryCaster() *types.ScriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}
