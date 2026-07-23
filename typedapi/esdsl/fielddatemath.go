package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldDateMath struct {
	v types.FieldDateMath
}

func NewFieldDateMath() *_fieldDateMath { _ = "STUB: not implemented"; return nil }

func (u *_fieldDateMath) DateMath(datemath string) *_fieldDateMath {
	_ = "STUB: not implemented"
	return nil
}

func (u *_fieldDateMath) Int64(int64 int64) *_fieldDateMath { _ = "STUB: not implemented"; return nil }

func (u *_fieldDateMath) FieldDateMathCaster() *types.FieldDateMath {
	_ = "STUB: not implemented"
	return nil
}
