package fieldsortnumerictype

type FieldSortNumericType struct {
	Name string
}

var (
	Long = FieldSortNumericType{"long"}

	Double = FieldSortNumericType{"double"}

	Date = FieldSortNumericType{"date"}

	Datenanos = FieldSortNumericType{"date_nanos"}
)

func (f FieldSortNumericType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FieldSortNumericType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldSortNumericType) String() string { _ = "STUB: not implemented"; return "" }
