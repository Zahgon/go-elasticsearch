package calendarinterval

type CalendarInterval struct {
	Name string
}

var (
	Second = CalendarInterval{"second"}

	Minute = CalendarInterval{"minute"}

	Hour = CalendarInterval{"hour"}

	Day = CalendarInterval{"day"}

	Week = CalendarInterval{"week"}

	Month = CalendarInterval{"month"}

	Quarter = CalendarInterval{"quarter"}

	Year = CalendarInterval{"year"}
)

func (c CalendarInterval) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CalendarInterval) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CalendarInterval) String() string { _ = "STUB: not implemented"; return "" }
