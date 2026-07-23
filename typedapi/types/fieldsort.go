package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type FieldSort struct {
	Format       *string                                    `json:"format,omitempty"`
	Missing      Missing                                    `json:"missing,omitempty"`
	Mode         *sortmode.SortMode                         `json:"mode,omitempty"`
	Nested       *NestedSortValue                           `json:"nested,omitempty"`
	NumericType  *fieldsortnumerictype.FieldSortNumericType `json:"numeric_type,omitempty"`
	Order        *sortorder.SortOrder                       `json:"order,omitempty"`
	UnmappedType *fieldtype.FieldType                       `json:"unmapped_type,omitempty"`
}

func (s *FieldSort) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldSort() *FieldSort { _ = "STUB: not implemented"; return nil }

type FieldSortVariant interface {
	FieldSortCaster() *FieldSort
}

func (s *FieldSort) FieldSortCaster() *FieldSort { _ = "STUB: not implemented"; return nil }
