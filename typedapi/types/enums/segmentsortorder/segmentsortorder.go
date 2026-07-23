package segmentsortorder

type SegmentSortOrder struct {
	Name string
}

var (
	Asc = SegmentSortOrder{"asc"}

	Desc = SegmentSortOrder{"desc"}
)

func (s SegmentSortOrder) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SegmentSortOrder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SegmentSortOrder) String() string { _ = "STUB: not implemented"; return "" }
