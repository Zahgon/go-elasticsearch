package snapshotsort

type SnapshotSort struct {
	Name string
}

var (
	Starttime = SnapshotSort{"start_time"}

	Duration = SnapshotSort{"duration"}

	Name = SnapshotSort{"name"}

	Indexcount = SnapshotSort{"index_count"}

	Repository = SnapshotSort{"repository"}

	Shardcount = SnapshotSort{"shard_count"}

	Failedshardcount = SnapshotSort{"failed_shard_count"}
)

func (s SnapshotSort) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SnapshotSort) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SnapshotSort) String() string { _ = "STUB: not implemented"; return "" }
