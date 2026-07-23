package queryruletype

type QueryRuleType struct {
	Name string
}

var (
	Pinned = QueryRuleType{"pinned"}

	Exclude = QueryRuleType{"exclude"}
)

func (q QueryRuleType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *QueryRuleType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (q QueryRuleType) String() string { _ = "STUB: not implemented"; return "" }
