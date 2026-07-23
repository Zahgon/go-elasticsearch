package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sizeField struct {
	v *types.SizeField
}

func NewSizeField(enabled bool) *_sizeField { _ = "STUB: not implemented"; return nil }

func (s *_sizeField) Enabled(enabled bool) *_sizeField { _ = "STUB: not implemented"; return nil }

func (s *_sizeField) SizeFieldCaster() *types.SizeField { _ = "STUB: not implemented"; return nil }
