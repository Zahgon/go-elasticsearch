package day

type Day struct {
	Name string
}

var (
	Sunday = Day{"sunday"}

	Monday = Day{"monday"}

	Tuesday = Day{"tuesday"}

	Wednesday = Day{"wednesday"}

	Thursday = Day{"thursday"}

	Friday = Day{"friday"}

	Saturday = Day{"saturday"}
)

func (d Day) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Day) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d Day) String() string { _ = "STUB: not implemented"; return "" }
