package rangerelation

type RangeRelation struct {
	Name string
}

var (
	Within = RangeRelation{"within"}

	Contains = RangeRelation{"contains"}

	Intersects = RangeRelation{"intersects"}
)

func (r RangeRelation) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RangeRelation) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RangeRelation) String() string { _ = "STUB: not implemented"; return "" }
