package managedby

type ManagedBy struct {
	Name string
}

var (
	Ilm = ManagedBy{"Index Lifecycle Management"}

	Datastream = ManagedBy{"Data stream lifecycle"}

	Unmanaged = ManagedBy{"Unmanaged"}
)

func (m ManagedBy) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ManagedBy) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m ManagedBy) String() string { _ = "STUB: not implemented"; return "" }
