package remoteclusterprivilege

type RemoteClusterPrivilege struct {
	Name string
}

var (
	Monitorenrich = RemoteClusterPrivilege{"monitor_enrich"}

	Monitorstats = RemoteClusterPrivilege{"monitor_stats"}
)

func (r RemoteClusterPrivilege) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RemoteClusterPrivilege) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r RemoteClusterPrivilege) String() string { _ = "STUB: not implemented"; return "" }
