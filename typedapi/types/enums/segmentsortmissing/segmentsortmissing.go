package segmentsortmissing

type SegmentSortMissing struct {
	Name string
}

var (
	Last = SegmentSortMissing{"_last"}

	First = SegmentSortMissing{"_first"}
)

func (s SegmentSortMissing) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SegmentSortMissing) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SegmentSortMissing) String() string { _ = "STUB: not implemented"; return "" }
