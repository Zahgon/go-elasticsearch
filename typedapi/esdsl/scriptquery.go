package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptQuery struct {
	v *types.ScriptQuery
}

func NewScriptQuery(script types.ScriptVariant) *_scriptQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptQuery) Script(script types.ScriptVariant) *_scriptQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptQuery) Boost(boost float32) *_scriptQuery { _ = "STUB: not implemented"; return nil }

func (s *_scriptQuery) QueryName_(queryname_ string) *_scriptQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_scriptQuery) ScriptQueryCaster() *types.ScriptQuery {
	_ = "STUB: not implemented"
	return nil
}
