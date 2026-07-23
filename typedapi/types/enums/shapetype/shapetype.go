package shapetype

type ShapeType struct {
	Name string
}

var (
	Geoshape = ShapeType{"geo_shape"}

	Shape = ShapeType{"shape"}
)

func (s ShapeType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShapeType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShapeType) String() string { _ = "STUB: not implemented"; return "" }
