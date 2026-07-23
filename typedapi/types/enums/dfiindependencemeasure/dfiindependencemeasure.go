package dfiindependencemeasure

type DFIIndependenceMeasure struct {
	Name string
}

var (
	Standardized = DFIIndependenceMeasure{"standardized"}

	Saturated = DFIIndependenceMeasure{"saturated"}

	Chisquared = DFIIndependenceMeasure{"chisquared"}
)

func (d DFIIndependenceMeasure) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DFIIndependenceMeasure) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DFIIndependenceMeasure) String() string { _ = "STUB: not implemented"; return "" }
