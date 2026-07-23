package catnodeattrscolumn

type CatNodeattrsColumn struct {
	Name string
}

var (
	Node = CatNodeattrsColumn{"node"}

	Id = CatNodeattrsColumn{"id"}

	Pid = CatNodeattrsColumn{"pid"}

	Host = CatNodeattrsColumn{"host"}

	Ip = CatNodeattrsColumn{"ip"}

	Port = CatNodeattrsColumn{"port"}

	Attr = CatNodeattrsColumn{"attr"}

	Value = CatNodeattrsColumn{"value"}
)

func (c CatNodeattrsColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatNodeattrsColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatNodeattrsColumn) String() string { _ = "STUB: not implemented"; return "" }
