package cattemplatescolumn

type CatTemplatesColumn struct {
	Name string
}

var (
	Name = CatTemplatesColumn{"name"}

	Indexpatterns = CatTemplatesColumn{"index_patterns"}

	Order = CatTemplatesColumn{"order"}

	Version = CatTemplatesColumn{"version"}

	Composedof = CatTemplatesColumn{"composed_of"}
)

func (c CatTemplatesColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatTemplatesColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatTemplatesColumn) String() string { _ = "STUB: not implemented"; return "" }
