package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _emptyObject struct {
	v *types.EmptyObject
}

func NewEmptyObject() *_emptyObject { _ = "STUB: not implemented"; return nil }

func (s *_emptyObject) EmptyObjectCaster() *types.EmptyObject {
	_ = "STUB: not implemented"
	return nil
}
