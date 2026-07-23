package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankContainer struct {
	v *types.RankContainer
}

func NewRankContainer() *_rankContainer { _ = "STUB: not implemented"; return nil }

func (s *_rankContainer) Rrf(rrf types.RrfRankVariant) *_rankContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankContainer) RankContainerCaster() *types.RankContainer {
	_ = "STUB: not implemented"
	return nil
}
