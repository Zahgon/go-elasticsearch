package memorystatus

type MemoryStatus struct {
	Name string
}

var (
	Ok = MemoryStatus{"ok"}

	Softlimit = MemoryStatus{"soft_limit"}

	Hardlimit = MemoryStatus{"hard_limit"}
)

func (m MemoryStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MemoryStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MemoryStatus) String() string { _ = "STUB: not implemented"; return "" }
