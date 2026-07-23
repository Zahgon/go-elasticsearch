package catdatafeedcolumn

type CatDatafeedColumn struct {
	Name string
}

var (
	Ae = CatDatafeedColumn{"ae"}

	Bc = CatDatafeedColumn{"bc"}

	Id = CatDatafeedColumn{"id"}

	Na = CatDatafeedColumn{"na"}

	Ne = CatDatafeedColumn{"ne"}

	Ni = CatDatafeedColumn{"ni"}

	Nn = CatDatafeedColumn{"nn"}

	Sba = CatDatafeedColumn{"sba"}

	Sc = CatDatafeedColumn{"sc"}

	Seah = CatDatafeedColumn{"seah"}

	St = CatDatafeedColumn{"st"}

	S = CatDatafeedColumn{"s"}
)

func (c CatDatafeedColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatDatafeedColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatDatafeedColumn) String() string { _ = "STUB: not implemented"; return "" }
