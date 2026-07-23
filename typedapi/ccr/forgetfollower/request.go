package forgetfollower

type Request struct {
	FollowerCluster     *string `json:"follower_cluster,omitempty"`
	FollowerIndex       *string `json:"follower_index,omitempty"`
	FollowerIndexUuid   *string `json:"follower_index_uuid,omitempty"`
	LeaderRemoteCluster *string `json:"leader_remote_cluster,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
