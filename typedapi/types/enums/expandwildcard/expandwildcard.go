package expandwildcard

type ExpandWildcard struct {
	Name string
}

var (
	All = ExpandWildcard{"all"}

	Open = ExpandWildcard{"open"}

	Closed = ExpandWildcard{"closed"}

	Hidden = ExpandWildcard{"hidden"}

	None = ExpandWildcard{"none"}
)

func (e ExpandWildcard) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ExpandWildcard) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e ExpandWildcard) String() string { _ = "STUB: not implemented"; return "" }
