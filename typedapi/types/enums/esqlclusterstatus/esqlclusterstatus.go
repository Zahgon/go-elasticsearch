package esqlclusterstatus

type EsqlClusterStatus struct {
	Name string
}

var (
	Running = EsqlClusterStatus{"running"}

	Successful = EsqlClusterStatus{"successful"}

	Partial = EsqlClusterStatus{"partial"}

	Skipped = EsqlClusterStatus{"skipped"}

	Failed = EsqlClusterStatus{"failed"}
)

func (e EsqlClusterStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EsqlClusterStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e EsqlClusterStatus) String() string { _ = "STUB: not implemented"; return "" }
