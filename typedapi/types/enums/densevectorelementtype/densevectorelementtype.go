package densevectorelementtype

type DenseVectorElementType struct {
	Name string
}

var (
	Bit = DenseVectorElementType{"bit"}

	Byte = DenseVectorElementType{"byte"}

	Float = DenseVectorElementType{"float"}

	Bfloat16 = DenseVectorElementType{"bfloat16"}
)

func (d DenseVectorElementType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DenseVectorElementType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DenseVectorElementType) String() string { _ = "STUB: not implemented"; return "" }
