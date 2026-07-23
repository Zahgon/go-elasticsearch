package valuetype

type ValueType struct {
	Name string
}

var (
	String = ValueType{"string"}

	Long = ValueType{"long"}

	Double = ValueType{"double"}

	Number = ValueType{"number"}

	Date = ValueType{"date"}

	Datenanos = ValueType{"date_nanos"}

	Ip = ValueType{"ip"}

	Numeric = ValueType{"numeric"}

	Geopoint = ValueType{"geo_point"}

	Boolean = ValueType{"boolean"}
)

func (v ValueType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *ValueType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (v ValueType) String() string { _ = "STUB: not implemented"; return "" }
