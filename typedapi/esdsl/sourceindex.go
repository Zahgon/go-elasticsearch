package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sourceIndex struct {
	v *types.SourceIndex
}

func NewSourceIndex() *_sourceIndex { _ = "STUB: not implemented"; return nil }

func (s *_sourceIndex) Index(indexname string) *_sourceIndex { _ = "STUB: not implemented"; return nil }

func (s *_sourceIndex) SourceIndexCaster() *types.SourceIndex {
	_ = "STUB: not implemented"
	return nil
}
