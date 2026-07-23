package suggestsort

type SuggestSort struct {
	Name string
}

var (
	Score = SuggestSort{"score"}

	Frequency = SuggestSort{"frequency"}
)

func (s SuggestSort) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SuggestSort) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SuggestSort) String() string { _ = "STUB: not implemented"; return "" }
