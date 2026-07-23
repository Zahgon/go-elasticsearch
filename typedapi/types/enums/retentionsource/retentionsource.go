package retentionsource

type RetentionSource struct {
	Name string
}

var (
	Datastreamconfiguration = RetentionSource{"data_stream_configuration"}

	Defaultglobalretention = RetentionSource{"default_global_retention"}

	Maxglobalretention = RetentionSource{"max_global_retention"}

	Defaultfailuresretention = RetentionSource{"default_failures_retention"}
)

func (r RetentionSource) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RetentionSource) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RetentionSource) String() string { _ = "STUB: not implemented"; return "" }
