package connectorfieldtype

type ConnectorFieldType struct {
	Name string
}

var (
	Str = ConnectorFieldType{"str"}

	Int = ConnectorFieldType{"int"}

	List = ConnectorFieldType{"list"}

	Bool = ConnectorFieldType{"bool"}
)

func (c ConnectorFieldType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnectorFieldType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c ConnectorFieldType) String() string { _ = "STUB: not implemented"; return "" }
