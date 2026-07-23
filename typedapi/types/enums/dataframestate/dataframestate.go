package dataframestate

type DataframeState struct {
	Name string
}

var (
	Started = DataframeState{"started"}

	Stopped = DataframeState{"stopped"}

	Starting = DataframeState{"starting"}

	Stopping = DataframeState{"stopping"}

	Failed = DataframeState{"failed"}
)

func (d DataframeState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DataframeState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DataframeState) String() string { _ = "STUB: not implemented"; return "" }
