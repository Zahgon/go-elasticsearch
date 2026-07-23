package catcomponentcolumn

type CatComponentColumn struct {
	Name string
}

var (
	Name = CatComponentColumn{"name"}

	Version = CatComponentColumn{"version"}

	Aliascount = CatComponentColumn{"alias_count"}

	Mappingcount = CatComponentColumn{"mapping_count"}

	Settingscount = CatComponentColumn{"settings_count"}

	Metadatacount = CatComponentColumn{"metadata_count"}

	Includedin = CatComponentColumn{"included_in"}
)

func (c CatComponentColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatComponentColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatComponentColumn) String() string { _ = "STUB: not implemented"; return "" }
