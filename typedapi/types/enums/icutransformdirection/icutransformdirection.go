package icutransformdirection

type IcuTransformDirection struct {
	Name string
}

var (
	Forward = IcuTransformDirection{"forward"}

	Reverse = IcuTransformDirection{"reverse"}
)

func (i IcuTransformDirection) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuTransformDirection) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuTransformDirection) String() string { _ = "STUB: not implemented"; return "" }
