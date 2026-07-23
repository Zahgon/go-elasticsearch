package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _alias struct {
	v *types.Alias
}

func NewAlias() *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) Filter(filter types.QueryVariant) *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) IndexRouting(indexrouting string) *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) IsHidden(ishidden bool) *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) IsWriteIndex(iswriteindex bool) *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) Routing(routing string) *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) SearchRouting(searchrouting string) *_alias { _ = "STUB: not implemented"; return nil }

func (s *_alias) AliasCaster() *types.Alias { _ = "STUB: not implemented"; return nil }
