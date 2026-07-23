package amazonsagemakerelementtype

type AmazonSageMakerElementType struct {
	Name string
}

var (
	Byte = AmazonSageMakerElementType{"byte"}

	Float = AmazonSageMakerElementType{"float"}

	Bit = AmazonSageMakerElementType{"bit"}
)

func (a AmazonSageMakerElementType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AmazonSageMakerElementType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AmazonSageMakerElementType) String() string { _ = "STUB: not implemented"; return "" }
