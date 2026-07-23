package syncstatus

type SyncStatus struct {
	Name string
}

var (
	Canceling = SyncStatus{"canceling"}

	Canceled = SyncStatus{"canceled"}

	Completed = SyncStatus{"completed"}

	Error = SyncStatus{"error"}

	Inprogress = SyncStatus{"in_progress"}

	Pending = SyncStatus{"pending"}

	Suspended = SyncStatus{"suspended"}
)

func (s SyncStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SyncStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SyncStatus) String() string { _ = "STUB: not implemented"; return "" }
