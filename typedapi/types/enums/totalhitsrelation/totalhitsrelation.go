package totalhitsrelation

type TotalHitsRelation struct {
	Name string
}

var (
	Eq = TotalHitsRelation{"eq"}

	Gte = TotalHitsRelation{"gte"}
)

func (t TotalHitsRelation) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TotalHitsRelation) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TotalHitsRelation) String() string { _ = "STUB: not implemented"; return "" }
