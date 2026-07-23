package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pinnedDoc struct {
	v *types.PinnedDoc
}

func NewPinnedDoc() *_pinnedDoc { _ = "STUB: not implemented"; return nil }

func (s *_pinnedDoc) Id_(id string) *_pinnedDoc { _ = "STUB: not implemented"; return nil }

func (s *_pinnedDoc) Index_(indexname string) *_pinnedDoc { _ = "STUB: not implemented"; return nil }

func (s *_pinnedDoc) PinnedQueryCaster() *types.PinnedQuery { _ = "STUB: not implemented"; return nil }

func (s *_pinnedDoc) PinnedDocCaster() *types.PinnedDoc { _ = "STUB: not implemented"; return nil }
