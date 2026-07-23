package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scalarValue struct {
	v types.ScalarValue
}

func NewScalarValue() *_scalarValue { _ = "STUB: not implemented"; return nil }

func (u *_scalarValue) Int64(int64 int64) *_scalarValue { _ = "STUB: not implemented"; return nil }

func (u *_scalarValue) Float64(float64 types.Float64) *_scalarValue {
	_ = "STUB: not implemented"
	return nil
}

func (u *_scalarValue) String(string string) *_scalarValue { _ = "STUB: not implemented"; return nil }

func (u *_scalarValue) Bool(bool bool) *_scalarValue { _ = "STUB: not implemented"; return nil }

func (u *_scalarValue) Nil() *_scalarValue { _ = "STUB: not implemented"; return nil }

func (u *_scalarValue) ScalarValueCaster() *types.ScalarValue {
	_ = "STUB: not implemented"
	return nil
}
