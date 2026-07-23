package rankvectorelementtype

type RankVectorElementType struct {
	Name string
}

var (
	Byte = RankVectorElementType{"byte"}

	Float = RankVectorElementType{"float"}

	Bit = RankVectorElementType{"bit"}
)

func (r RankVectorElementType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RankVectorElementType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r RankVectorElementType) String() string { _ = "STUB: not implemented"; return "" }
