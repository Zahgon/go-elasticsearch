package jobblockedreason

type JobBlockedReason struct {
	Name string
}

var (
	Delete = JobBlockedReason{"delete"}

	Reset = JobBlockedReason{"reset"}

	Revert = JobBlockedReason{"revert"}
)

func (j JobBlockedReason) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobBlockedReason) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (j JobBlockedReason) String() string { _ = "STUB: not implemented"; return "" }
