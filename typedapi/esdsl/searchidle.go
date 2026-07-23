package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchIdle struct {
	v *types.SearchIdle
}

func NewSearchIdle() *_searchIdle { _ = "STUB: not implemented"; return nil }

func (s *_searchIdle) After(duration types.DurationVariant) *_searchIdle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchIdle) SearchIdleCaster() *types.SearchIdle { _ = "STUB: not implemented"; return nil }
