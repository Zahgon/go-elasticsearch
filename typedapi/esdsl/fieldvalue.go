package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldValue struct {
	v types.FieldValue
}

func NewFieldValue() *_fieldValue { _ = "STUB: not implemented"; return nil }

func (u *_fieldValue) Int64(int64 int64) *_fieldValue { _ = "STUB: not implemented"; return nil }

func (u *_fieldValue) Float64(float64 types.Float64) *_fieldValue {
	_ = "STUB: not implemented"
	return nil
}

func (u *_fieldValue) String(string string) *_fieldValue { _ = "STUB: not implemented"; return nil }

func (u *_fieldValue) Bool(bool bool) *_fieldValue { _ = "STUB: not implemented"; return nil }

func (u *_fieldValue) Nil() *_fieldValue { _ = "STUB: not implemented"; return nil }

func (u *_fieldValue) FieldValueCaster() *types.FieldValue { _ = "STUB: not implemented"; return nil }
