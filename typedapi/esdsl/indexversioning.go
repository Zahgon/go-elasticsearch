package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexVersioning struct {
	v *types.IndexVersioning
}

func NewIndexVersioning() *_indexVersioning { _ = "STUB: not implemented"; return nil }

func (s *_indexVersioning) Created(versionstring string) *_indexVersioning {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexVersioning) CreatedString(createdstring string) *_indexVersioning {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexVersioning) IndexVersioningCaster() *types.IndexVersioning {
	_ = "STUB: not implemented"
	return nil
}
