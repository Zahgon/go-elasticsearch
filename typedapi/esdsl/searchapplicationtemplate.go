package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchApplicationTemplate struct {
	v *types.SearchApplicationTemplate
}

func NewSearchApplicationTemplate(script types.ScriptVariant) *_searchApplicationTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchApplicationTemplate) Script(script types.ScriptVariant) *_searchApplicationTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchApplicationTemplate) SearchApplicationTemplateCaster() *types.SearchApplicationTemplate {
	_ = "STUB: not implemented"
	return nil
}
