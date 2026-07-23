package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _multiTermLookup struct {
	v *types.MultiTermLookup
}

func NewMultiTermLookup() *_multiTermLookup { _ = "STUB: not implemented"; return nil }

func (s *_multiTermLookup) Field(field string) *_multiTermLookup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermLookup) Missing(missing types.MissingVariant) *_multiTermLookup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermLookup) Script(script types.ScriptVariant) *_multiTermLookup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiTermLookup) MultiTermLookupCaster() *types.MultiTermLookup {
	_ = "STUB: not implemented"
	return nil
}
