package types

type Status struct {
	IncludeGlobalState bool                          `json:"include_global_state"`
	Indices            map[string]SnapshotIndexStats `json:"indices"`

	Repository string `json:"repository"`

	ShardsStats SnapshotShardsStats `json:"shards_stats"`

	Snapshot string `json:"snapshot"`

	State string `json:"state"`

	Stats SnapshotStats `json:"stats"`

	Uuid string `json:"uuid"`
}

func (s *Status) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStatus() *Status { _ = "STUB: not implemented"; return nil }
