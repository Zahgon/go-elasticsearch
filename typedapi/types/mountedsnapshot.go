package types

type MountedSnapshot struct {
	Indices  []string        `json:"indices"`
	Shards   ShardStatistics `json:"shards"`
	Snapshot string          `json:"snapshot"`
}

func (s *MountedSnapshot) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMountedSnapshot() *MountedSnapshot { _ = "STUB: not implemented"; return nil }
