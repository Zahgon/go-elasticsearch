package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _addAction struct {
	v *types.AddAction
}

func NewAddAction() *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) Alias(indexalias string) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) Aliases(aliases ...string) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) Filter(filter types.QueryVariant) *_addAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_addAction) Index(indexname string) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) IndexRouting(indexrouting string) *_addAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_addAction) Indices(indices ...string) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) IsHidden(ishidden bool) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) IsWriteIndex(iswriteindex bool) *_addAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_addAction) MustExist(mustexist bool) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) Routing(routing string) *_addAction { _ = "STUB: not implemented"; return nil }

func (s *_addAction) SearchRouting(searchrouting string) *_addAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_addAction) IndicesActionCaster() *types.IndicesAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_addAction) AddActionCaster() *types.AddAction { _ = "STUB: not implemented"; return nil }
