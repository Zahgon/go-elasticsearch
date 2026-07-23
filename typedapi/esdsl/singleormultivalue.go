package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _singleOrMultiValue struct {
	v types.SingleOrMultiValue
}

func NewSingleOrMultiValue() *_singleOrMultiValue { _ = "STUB: not implemented"; return nil }

func (u *_singleOrMultiValue) SingleOrMultiValueCaster() *types.SingleOrMultiValue {
	_ = "STUB: not implemented"
	return nil
}
