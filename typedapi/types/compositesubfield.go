package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/runtimefieldtype"
)

type CompositeSubField struct {
	Type runtimefieldtype.RuntimeFieldType `json:"type"`
}

func NewCompositeSubField() *CompositeSubField { _ = "STUB: not implemented"; return nil }

type CompositeSubFieldVariant interface {
	CompositeSubFieldCaster() *CompositeSubField
}

func (s *CompositeSubField) CompositeSubFieldCaster() *CompositeSubField {
	_ = "STUB: not implemented"
	return nil
}
