package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _latest struct {
	v *types.Latest
}

func NewLatest() *_latest { _ = "STUB: not implemented"; return nil }

func (s *_latest) Sort(field string) *_latest { _ = "STUB: not implemented"; return nil }

func (s *_latest) UniqueKey(uniquekeys ...string) *_latest { _ = "STUB: not implemented"; return nil }

func (s *_latest) LatestCaster() *types.Latest { _ = "STUB: not implemented"; return nil }
