package delimitedpayloadencoding

type DelimitedPayloadEncoding struct {
	Name string
}

var (
	Int = DelimitedPayloadEncoding{"int"}

	Float = DelimitedPayloadEncoding{"float"}

	Identity = DelimitedPayloadEncoding{"identity"}
)

func (d DelimitedPayloadEncoding) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DelimitedPayloadEncoding) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DelimitedPayloadEncoding) String() string { _ = "STUB: not implemented"; return "" }
