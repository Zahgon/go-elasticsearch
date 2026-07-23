package catpluginscolumn

type CatPluginsColumn struct {
	Name string
}

var (
	Id = CatPluginsColumn{"id"}

	Name = CatPluginsColumn{"name"}

	Component = CatPluginsColumn{"component"}

	Version = CatPluginsColumn{"version"}

	Description = CatPluginsColumn{"description"}
)

func (c CatPluginsColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatPluginsColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatPluginsColumn) String() string { _ = "STUB: not implemented"; return "" }
