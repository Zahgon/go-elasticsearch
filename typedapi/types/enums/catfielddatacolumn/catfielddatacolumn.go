package catfielddatacolumn

type CatFieldDataColumn struct {
	Name string
}

var (
	Id = CatFieldDataColumn{"id"}

	Host = CatFieldDataColumn{"host"}

	Ip = CatFieldDataColumn{"ip"}

	Node = CatFieldDataColumn{"node"}

	Field = CatFieldDataColumn{"field"}

	Size = CatFieldDataColumn{"size"}
)

func (c CatFieldDataColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatFieldDataColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatFieldDataColumn) String() string { _ = "STUB: not implemented"; return "" }
