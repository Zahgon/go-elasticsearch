package cohereembeddingtype

type CohereEmbeddingType struct {
	Name string
}

var (
	Binary = CohereEmbeddingType{"binary"}

	Bit = CohereEmbeddingType{"bit"}

	Byte = CohereEmbeddingType{"byte"}

	Float = CohereEmbeddingType{"float"}

	Int8 = CohereEmbeddingType{"int8"}
)

func (c CohereEmbeddingType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CohereEmbeddingType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CohereEmbeddingType) String() string { _ = "STUB: not implemented"; return "" }
