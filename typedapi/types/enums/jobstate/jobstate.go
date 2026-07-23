package jobstate

type JobState struct {
	Name string
}

var (
	Closing = JobState{"closing"}

	Closed = JobState{"closed"}

	Opened = JobState{"opened"}

	Failed = JobState{"failed"}

	Opening = JobState{"opening"}
)

func (j JobState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (j JobState) String() string { _ = "STUB: not implemented"; return "" }
