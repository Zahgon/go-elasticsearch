package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _namedValue struct {
	v types.NamedValue
}

func NewNamedValue(namedvalue map[string][]types.FieldValue) *_namedValue {
	_ = "STUB: not implemented"
	return nil
}

func (u *_namedValue) NamedValueCaster() *types.NamedValue { _ = "STUB: not implemented"; return nil }
