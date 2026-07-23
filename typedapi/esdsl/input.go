package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _input struct {
	v *types.Input
}

func NewInput() *_input { _ = "STUB: not implemented"; return nil }

func (s *_input) FieldNames(names ...string) *_input { _ = "STUB: not implemented"; return nil }

func (s *_input) InputCaster() *types.Input { _ = "STUB: not implemented"; return nil }
