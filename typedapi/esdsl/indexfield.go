package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexField struct {
	v *types.IndexField
}

func NewIndexField(enabled bool) *_indexField { _ = "STUB: not implemented"; return nil }

func (s *_indexField) Enabled(enabled bool) *_indexField { _ = "STUB: not implemented"; return nil }

func (s *_indexField) IndexFieldCaster() *types.IndexField { _ = "STUB: not implemented"; return nil }
