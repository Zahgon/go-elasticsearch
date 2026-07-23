package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _charFilter struct {
	v types.CharFilter
}

func NewCharFilter() *_charFilter { _ = "STUB: not implemented"; return nil }

func (u *_charFilter) String(string string) *_charFilter { _ = "STUB: not implemented"; return nil }

func (u *_charFilter) CharFilterDefinition(charfilterdefinition types.CharFilterDefinitionVariant) *_charFilter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) CharFilterCaster() *types.CharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilter) CharFilterCaster() *types.CharFilter { _ = "STUB: not implemented"; return nil }
