package syncjobtype

type SyncJobType struct {
	Name string
}

var (
	Full = SyncJobType{"full"}

	Incremental = SyncJobType{"incremental"}

	Accesscontrol = SyncJobType{"access_control"}
)

func (s SyncJobType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SyncJobType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SyncJobType) String() string { _ = "STUB: not implemented"; return "" }
