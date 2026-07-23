package datafeedstate

type DatafeedState struct {
	Name string
}

var (
	Started = DatafeedState{"started"}

	Stopped = DatafeedState{"stopped"}

	Starting = DatafeedState{"starting"}

	Stopping = DatafeedState{"stopping"}
)

func (d DatafeedState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DatafeedState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DatafeedState) String() string { _ = "STUB: not implemented"; return "" }
