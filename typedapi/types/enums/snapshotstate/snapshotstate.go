package snapshotstate

type SnapshotState struct {
	Name string
}

var (
	INPROGRESS = SnapshotState{"IN_PROGRESS"}

	SUCCESS = SnapshotState{"SUCCESS"}

	FAILED = SnapshotState{"FAILED"}

	PARTIAL = SnapshotState{"PARTIAL"}

	INCOMPATIBLE = SnapshotState{"INCOMPATIBLE"}
)

func (s SnapshotState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SnapshotState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SnapshotState) String() string { _ = "STUB: not implemented"; return "" }
