package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/followerindexstatus"
)

type FollowerIndex struct {
	FollowerIndex string `json:"follower_index"`

	LeaderIndex string `json:"leader_index"`

	Parameters *FollowerIndexParameters `json:"parameters,omitempty"`

	RemoteCluster string `json:"remote_cluster"`

	Status followerindexstatus.FollowerIndexStatus `json:"status"`
}

func (s *FollowerIndex) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFollowerIndex() *FollowerIndex { _ = "STUB: not implemented"; return nil }
