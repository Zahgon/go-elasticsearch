package timeunit

type TimeUnit struct {
	Name string
}

var (
	Nanoseconds = TimeUnit{"nanos"}

	Microseconds = TimeUnit{"micros"}

	Milliseconds = TimeUnit{"ms"}

	Seconds = TimeUnit{"s"}

	Minutes = TimeUnit{"m"}

	Hours = TimeUnit{"h"}

	Days = TimeUnit{"d"}
)

func (t TimeUnit) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TimeUnit) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TimeUnit) String() string { _ = "STUB: not implemented"; return "" }
