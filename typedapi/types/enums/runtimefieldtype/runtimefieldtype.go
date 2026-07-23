package runtimefieldtype

type RuntimeFieldType struct {
	Name string
}

var (
	Boolean = RuntimeFieldType{"boolean"}

	Composite = RuntimeFieldType{"composite"}

	Date = RuntimeFieldType{"date"}

	Double = RuntimeFieldType{"double"}

	Geopoint = RuntimeFieldType{"geo_point"}

	Geoshape = RuntimeFieldType{"geo_shape"}

	Ip = RuntimeFieldType{"ip"}

	Keyword = RuntimeFieldType{"keyword"}

	Long = RuntimeFieldType{"long"}

	Lookup = RuntimeFieldType{"lookup"}
)

func (r RuntimeFieldType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RuntimeFieldType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RuntimeFieldType) String() string { _ = "STUB: not implemented"; return "" }
