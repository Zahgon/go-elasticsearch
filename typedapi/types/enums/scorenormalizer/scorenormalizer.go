package scorenormalizer

type ScoreNormalizer struct {
	Name string
}

var (
	None = ScoreNormalizer{"none"}

	Minmax = ScoreNormalizer{"minmax"}

	L2norm = ScoreNormalizer{"l2_norm"}
)

func (s ScoreNormalizer) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ScoreNormalizer) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ScoreNormalizer) String() string { _ = "STUB: not implemented"; return "" }
