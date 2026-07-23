package dynamicmapping

type DynamicMapping struct {
	Name string
}

var (
	Strict = DynamicMapping{"strict"}

	Runtime = DynamicMapping{"runtime"}

	True = DynamicMapping{"true"}

	False = DynamicMapping{"false"}
)

func (d *DynamicMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (d DynamicMapping) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DynamicMapping) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DynamicMapping) String() string { _ = "STUB: not implemented"; return "" }
