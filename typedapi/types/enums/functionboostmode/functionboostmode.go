package functionboostmode

type FunctionBoostMode struct {
	Name string
}

var (
	Multiply = FunctionBoostMode{"multiply"}

	Replace = FunctionBoostMode{"replace"}

	Sum = FunctionBoostMode{"sum"}

	Avg = FunctionBoostMode{"avg"}

	Max = FunctionBoostMode{"max"}

	Min = FunctionBoostMode{"min"}
)

func (f FunctionBoostMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FunctionBoostMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FunctionBoostMode) String() string { _ = "STUB: not implemented"; return "" }
