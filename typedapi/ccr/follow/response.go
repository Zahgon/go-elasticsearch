package follow

type Response struct {
	FollowIndexCreated     bool `json:"follow_index_created"`
	FollowIndexShardsAcked bool `json:"follow_index_shards_acked"`
	IndexFollowingStarted  bool `json:"index_following_started"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
