package queryrulecriteriatype

type QueryRuleCriteriaType struct {
	Name string
}

var (
	Global = QueryRuleCriteriaType{"global"}

	Exact = QueryRuleCriteriaType{"exact"}

	Fuzzy = QueryRuleCriteriaType{"fuzzy"}

	Prefix = QueryRuleCriteriaType{"prefix"}

	Suffix = QueryRuleCriteriaType{"suffix"}

	Contains = QueryRuleCriteriaType{"contains"}

	Lt = QueryRuleCriteriaType{"lt"}

	Lte = QueryRuleCriteriaType{"lte"}

	Gt = QueryRuleCriteriaType{"gt"}

	Gte = QueryRuleCriteriaType{"gte"}

	Always = QueryRuleCriteriaType{"always"}
)

func (q QueryRuleCriteriaType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *QueryRuleCriteriaType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (q QueryRuleCriteriaType) String() string { _ = "STUB: not implemented"; return "" }
