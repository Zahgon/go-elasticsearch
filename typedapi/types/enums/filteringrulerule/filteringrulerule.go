package filteringrulerule

type FilteringRuleRule struct {
	Name string
}

var (
	Contains = FilteringRuleRule{"contains"}

	Endswith = FilteringRuleRule{"ends_with"}

	Equals = FilteringRuleRule{"equals"}

	Regex = FilteringRuleRule{"regex"}

	Startswith = FilteringRuleRule{"starts_with"}

	Greaterthan = FilteringRuleRule{">"}

	Lessthan = FilteringRuleRule{"<"}
)

func (f FilteringRuleRule) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilteringRuleRule) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FilteringRuleRule) String() string { _ = "STUB: not implemented"; return "" }
