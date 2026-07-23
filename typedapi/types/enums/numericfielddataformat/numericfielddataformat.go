package numericfielddataformat

type NumericFielddataFormat struct {
	Name string
}

var (
	Array = NumericFielddataFormat{"array"}

	Disabled = NumericFielddataFormat{"disabled"}
)

func (n NumericFielddataFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NumericFielddataFormat) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n NumericFielddataFormat) String() string { _ = "STUB: not implemented"; return "" }
