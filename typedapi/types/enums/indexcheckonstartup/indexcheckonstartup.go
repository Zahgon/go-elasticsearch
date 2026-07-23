package indexcheckonstartup

type IndexCheckOnStartup struct {
	Name string
}

var (
	True = IndexCheckOnStartup{"true"}

	False = IndexCheckOnStartup{"false"}

	Checksum = IndexCheckOnStartup{"checksum"}
)

func (i *IndexCheckOnStartup) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndexCheckOnStartup) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexCheckOnStartup) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndexCheckOnStartup) String() string { _ = "STUB: not implemented"; return "" }
