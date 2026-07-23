package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _fieldSort struct {
	v *types.FieldSort
}

func NewFieldSort(order sortorder.SortOrder) *_fieldSort { _ = "STUB: not implemented"; return nil }

func (s *_fieldSort) Format(format string) *_fieldSort { _ = "STUB: not implemented"; return nil }

func (s *_fieldSort) Missing(missing types.MissingVariant) *_fieldSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSort) Mode(mode sortmode.SortMode) *_fieldSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSort) Nested(nested types.NestedSortValueVariant) *_fieldSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSort) NumericType(numerictype fieldsortnumerictype.FieldSortNumericType) *_fieldSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSort) Order(order sortorder.SortOrder) *_fieldSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSort) UnmappedType(unmappedtype fieldtype.FieldType) *_fieldSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSort) FieldSortCaster() *types.FieldSort { _ = "STUB: not implemented"; return nil }
