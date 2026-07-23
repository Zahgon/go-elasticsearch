package geovalidationmethod

type GeoValidationMethod struct {
	Name string
}

var (
	Coerce = GeoValidationMethod{"coerce"}

	Ignoremalformed = GeoValidationMethod{"ignore_malformed"}

	Strict = GeoValidationMethod{"strict"}
)

func (g GeoValidationMethod) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoValidationMethod) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GeoValidationMethod) String() string { _ = "STUB: not implemented"; return "" }
