package fieldvaluefactormodifier

type FieldValueFactorModifier struct {
	Name string
}

var (
	None = FieldValueFactorModifier{"none"}

	Log = FieldValueFactorModifier{"log"}

	Log1p = FieldValueFactorModifier{"log1p"}

	Log2p = FieldValueFactorModifier{"log2p"}

	Ln = FieldValueFactorModifier{"ln"}

	Ln1p = FieldValueFactorModifier{"ln1p"}

	Ln2p = FieldValueFactorModifier{"ln2p"}

	Square = FieldValueFactorModifier{"square"}

	Sqrt = FieldValueFactorModifier{"sqrt"}

	Reciprocal = FieldValueFactorModifier{"reciprocal"}
)

func (f FieldValueFactorModifier) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FieldValueFactorModifier) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldValueFactorModifier) String() string { _ = "STUB: not implemented"; return "" }
