package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardstate"
)

type RepositoryStatsShards struct {
	Complete   int                           `json:"complete"`
	Incomplete int                           `json:"incomplete"`
	States     map[shardstate.ShardState]int `json:"states"`
	Total      int                           `json:"total"`
}

func (s *RepositoryStatsShards) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoryStatsShards() *RepositoryStatsShards { _ = "STUB: not implemented"; return nil }
