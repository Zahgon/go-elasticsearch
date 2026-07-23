package resultposition

type ResultPosition struct {
	Name string
}

var (
	Tail = ResultPosition{"tail"}

	Head = ResultPosition{"head"}
)

func (r ResultPosition) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ResultPosition) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r ResultPosition) String() string { _ = "STUB: not implemented"; return "" }
