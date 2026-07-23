package snapshotupgradestate

type SnapshotUpgradeState struct {
	Name string
}

var (
	Loadingoldstate = SnapshotUpgradeState{"loading_old_state"}

	Savingnewstate = SnapshotUpgradeState{"saving_new_state"}

	Stopped = SnapshotUpgradeState{"stopped"}

	Failed = SnapshotUpgradeState{"failed"}
)

func (s SnapshotUpgradeState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SnapshotUpgradeState) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SnapshotUpgradeState) String() string { _ = "STUB: not implemented"; return "" }
