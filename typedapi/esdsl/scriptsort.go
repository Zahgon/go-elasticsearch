package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptsorttype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _scriptSort struct {
	v *types.ScriptSort
}

func NewScriptSort(script types.ScriptVariant) *_scriptSort { _ = "STUB: not implemented"; return nil }

func (s *_scriptSort) Mode(mode sortmode.SortMode) *_scriptSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptSort) Nested(nested types.NestedSortValueVariant) *_scriptSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptSort) Order(order sortorder.SortOrder) *_scriptSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptSort) Script(script types.ScriptVariant) *_scriptSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptSort) Type(type_ scriptsorttype.ScriptSortType) *_scriptSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptSort) SortOptionsCaster() *types.SortOptions { _ = "STUB: not implemented"; return nil }

func (s *_scriptSort) ScriptSortCaster() *types.ScriptSort { _ = "STUB: not implemented"; return nil }
