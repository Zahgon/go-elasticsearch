package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptSource struct {
	v types.ScriptSource
}

func NewScriptSource() *_scriptSource { _ = "STUB: not implemented"; return nil }

func (u *_scriptSource) String(string string) *_scriptSource { _ = "STUB: not implemented"; return nil }

func (u *_scriptSource) SearchRequestBody(searchrequestbody types.SearchRequestBodyVariant) *_scriptSource {
	_ = "STUB: not implemented"
	return nil
}

func (u *_searchRequestBody) ScriptSourceCaster() *types.ScriptSource {
	_ = "STUB: not implemented"
	return nil
}

func (u *_scriptSource) ScriptSourceCaster() *types.ScriptSource {
	_ = "STUB: not implemented"
	return nil
}
