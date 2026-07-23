package suggestmode

type SuggestMode struct {
	Name string
}

var (
	Missing = SuggestMode{"missing"}

	Popular = SuggestMode{"popular"}

	Always = SuggestMode{"always"}
)

func (s SuggestMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SuggestMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SuggestMode) String() string { _ = "STUB: not implemented"; return "" }
