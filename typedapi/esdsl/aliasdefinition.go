package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _aliasDefinition struct {
	v *types.AliasDefinition
}

func NewAliasDefinition() *_aliasDefinition { _ = "STUB: not implemented"; return nil }

func (s *_aliasDefinition) Filter(filter types.QueryVariant) *_aliasDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aliasDefinition) IndexRouting(indexrouting string) *_aliasDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aliasDefinition) IsHidden(ishidden bool) *_aliasDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aliasDefinition) IsWriteIndex(iswriteindex bool) *_aliasDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aliasDefinition) Routing(routing string) *_aliasDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aliasDefinition) SearchRouting(searchrouting string) *_aliasDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aliasDefinition) AliasDefinitionCaster() *types.AliasDefinition {
	_ = "STUB: not implemented"
	return nil
}
