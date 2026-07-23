package clustersearchstatus

type ClusterSearchStatus struct {
	Name string
}

var (
	Running = ClusterSearchStatus{"running"}

	Successful = ClusterSearchStatus{"successful"}

	Partial = ClusterSearchStatus{"partial"}

	Skipped = ClusterSearchStatus{"skipped"}

	Failed = ClusterSearchStatus{"failed"}
)

func (c ClusterSearchStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterSearchStatus) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c ClusterSearchStatus) String() string { _ = "STUB: not implemented"; return "" }
