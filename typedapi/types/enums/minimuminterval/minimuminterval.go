package minimuminterval

type MinimumInterval struct {
	Name string
}

var (
	Second = MinimumInterval{"second"}

	Minute = MinimumInterval{"minute"}

	Hour = MinimumInterval{"hour"}

	Day = MinimumInterval{"day"}

	Month = MinimumInterval{"month"}

	Year = MinimumInterval{"year"}
)

func (m MinimumInterval) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MinimumInterval) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MinimumInterval) String() string { _ = "STUB: not implemented"; return "" }
