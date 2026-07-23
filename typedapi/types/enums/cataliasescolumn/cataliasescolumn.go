package cataliasescolumn

type CatAliasesColumn struct {
	Name string
}

var (
	Alias = CatAliasesColumn{"alias"}

	Index = CatAliasesColumn{"index"}

	Filter = CatAliasesColumn{"filter"}

	Routingindex = CatAliasesColumn{"routing.index"}

	Routingsearch = CatAliasesColumn{"routing.search"}

	Iswriteindex = CatAliasesColumn{"is_write_index"}
)

func (c CatAliasesColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatAliasesColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatAliasesColumn) String() string { _ = "STUB: not implemented"; return "" }
