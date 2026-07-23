package types

import (
	"encoding/json"
)

type AsyncSearch struct {
	Aggregations map[string]Aggregate       `json:"aggregations,omitempty"`
	Clusters_    *ClusterStatistics         `json:"_clusters,omitempty"`
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Hits         HitsMetadata               `json:"hits"`
	MaxScore     *Float64                   `json:"max_score,omitempty"`

	NumReducePhases *int64   `json:"num_reduce_phases,omitempty"`
	PitId           *string  `json:"pit_id,omitempty"`
	Profile         *Profile `json:"profile,omitempty"`
	ScrollId_       *string  `json:"_scroll_id,omitempty"`

	Shards_         ShardStatistics      `json:"_shards"`
	Suggest         map[string][]Suggest `json:"suggest,omitempty"`
	TerminatedEarly *bool                `json:"terminated_early,omitempty"`
	TimedOut        bool                 `json:"timed_out"`
	Took            int64                `json:"took"`
}

func (s *AsyncSearch) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAsyncSearch() *AsyncSearch { _ = "STUB: not implemented"; return nil }
