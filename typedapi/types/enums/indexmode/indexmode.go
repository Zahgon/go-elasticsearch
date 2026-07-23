package indexmode

type IndexMode struct {
	Name string
}

var (
	Standard = IndexMode{"standard"}

	Timeseries = IndexMode{"time_series"}

	Logsdb = IndexMode{"logsdb"}

	Lookup = IndexMode{"lookup"}
)

func (i IndexMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i IndexMode) String() string { _ = "STUB: not implemented"; return "" }
