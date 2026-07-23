package distanceunit

type DistanceUnit struct {
	Name string
}

var (
	Inches = DistanceUnit{"in"}

	Feet = DistanceUnit{"ft"}

	Yards = DistanceUnit{"yd"}

	Miles = DistanceUnit{"mi"}

	Nauticmiles = DistanceUnit{"nmi"}

	Kilometers = DistanceUnit{"km"}

	Meters = DistanceUnit{"m"}

	Centimeters = DistanceUnit{"cm"}

	Millimeters = DistanceUnit{"mm"}
)

func (d DistanceUnit) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DistanceUnit) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DistanceUnit) String() string { _ = "STUB: not implemented"; return "" }
