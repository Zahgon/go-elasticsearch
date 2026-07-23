package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/runtimefieldtype"
)

type _compositeSubField struct {
	v *types.CompositeSubField
}

func NewCompositeSubField(type_ runtimefieldtype.RuntimeFieldType) *_compositeSubField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeSubField) Type(type_ runtimefieldtype.RuntimeFieldType) *_compositeSubField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeSubField) CompositeSubFieldCaster() *types.CompositeSubField {
	_ = "STUB: not implemented"
	return nil
}
