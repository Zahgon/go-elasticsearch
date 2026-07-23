package catcountcolumn

type CatCountColumn struct {
	Name string
}

var (
	Epoch = CatCountColumn{"epoch"}

	Timestamp = CatCountColumn{"timestamp"}

	Count = CatCountColumn{"count"}
)

func (c CatCountColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatCountColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatCountColumn) String() string { _ = "STUB: not implemented"; return "" }
