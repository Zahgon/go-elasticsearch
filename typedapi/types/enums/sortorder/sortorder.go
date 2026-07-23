package sortorder

type SortOrder struct {
	Name string
}

var (
	Asc = SortOrder{"asc"}

	Desc = SortOrder{"desc"}
)

func (s SortOrder) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SortOrder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SortOrder) String() string { _ = "STUB: not implemented"; return "" }
