package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rrfRank struct {
	v *types.RrfRank
}

func NewRrfRank() *_rrfRank { _ = "STUB: not implemented"; return nil }

func (s *_rrfRank) RankConstant(rankconstant int64) *_rrfRank {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rrfRank) RankWindowSize(rankwindowsize int64) *_rrfRank {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rrfRank) RankContainerCaster() *types.RankContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rrfRank) RrfRankCaster() *types.RrfRank { _ = "STUB: not implemented"; return nil }
