package converttype

type ConvertType struct {
	Name string
}

var (
	Integer = ConvertType{"integer"}

	Long = ConvertType{"long"}

	Double = ConvertType{"double"}

	Float = ConvertType{"float"}

	Boolean = ConvertType{"boolean"}

	Ip = ConvertType{"ip"}

	String = ConvertType{"string"}

	Auto = ConvertType{"auto"}
)

func (c ConvertType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConvertType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ConvertType) String() string { _ = "STUB: not implemented"; return "" }
