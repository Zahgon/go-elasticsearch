package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _suggestContext struct {
	v *types.SuggestContext
}

func NewSuggestContext(type_ string) *_suggestContext { _ = "STUB: not implemented"; return nil }

func (s *_suggestContext) Name(name string) *_suggestContext { _ = "STUB: not implemented"; return nil }

func (s *_suggestContext) Path(field string) *_suggestContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestContext) Precision(precision string) *_suggestContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestContext) Type(type_ string) *_suggestContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestContext) SuggestContextCaster() *types.SuggestContext {
	_ = "STUB: not implemented"
	return nil
}
