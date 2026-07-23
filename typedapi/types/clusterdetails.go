package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clustersearchstatus"
)

type ClusterDetails struct {
	Failures []ShardFailure                          `json:"failures,omitempty"`
	Indices  string                                  `json:"indices"`
	Shards_  *ShardStatistics                        `json:"_shards,omitempty"`
	Status   clustersearchstatus.ClusterSearchStatus `json:"status"`
	TimedOut bool                                    `json:"timed_out"`
	Took     *int64                                  `json:"took,omitempty"`
}

func (s *ClusterDetails) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterDetails() *ClusterDetails { _ = "STUB: not implemented"; return nil }
