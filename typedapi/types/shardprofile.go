package types

type ShardProfile struct {
	Aggregations []AggregationProfile `json:"aggregations"`
	Cluster      string               `json:"cluster"`
	Dfs          *DfsProfile          `json:"dfs,omitempty"`
	Fetch        *FetchProfile        `json:"fetch,omitempty"`
	Id           string               `json:"id"`
	Index        string               `json:"index"`
	NodeId       string               `json:"node_id"`
	Searches     []SearchProfile      `json:"searches"`
	ShardId      int                  `json:"shard_id"`
}

func (s *ShardProfile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardProfile() *ShardProfile { _ = "STUB: not implemented"; return nil }
