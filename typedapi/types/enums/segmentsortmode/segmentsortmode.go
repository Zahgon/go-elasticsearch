package segmentsortmode

type SegmentSortMode struct {
	Name string
}

var (
	Min = SegmentSortMode{"min"}

	Max = SegmentSortMode{"max"}
)

func (s SegmentSortMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SegmentSortMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SegmentSortMode) String() string { _ = "STUB: not implemented"; return "" }
