package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptsorttype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type ScriptSort struct {
	Mode   *sortmode.SortMode             `json:"mode,omitempty"`
	Nested *NestedSortValue               `json:"nested,omitempty"`
	Order  *sortorder.SortOrder           `json:"order,omitempty"`
	Script Script                         `json:"script"`
	Type   *scriptsorttype.ScriptSortType `json:"type,omitempty"`
}

func NewScriptSort() *ScriptSort { _ = "STUB: not implemented"; return nil }

type ScriptSortVariant interface {
	ScriptSortCaster() *ScriptSort
}

func (s *ScriptSort) ScriptSortCaster() *ScriptSort { _ = "STUB: not implemented"; return nil }
