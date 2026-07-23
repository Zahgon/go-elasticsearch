package catmastercolumn

type CatMasterColumn struct {
	Name string
}

var (
	Id = CatMasterColumn{"id"}

	Host = CatMasterColumn{"host"}

	Ip = CatMasterColumn{"ip"}

	Node = CatMasterColumn{"node"}
)

func (c CatMasterColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatMasterColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatMasterColumn) String() string { _ = "STUB: not implemented"; return "" }
