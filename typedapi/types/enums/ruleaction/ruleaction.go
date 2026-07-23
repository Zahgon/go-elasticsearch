package ruleaction

type RuleAction struct {
	Name string
}

var (
	Skipresult = RuleAction{"skip_result"}

	Skipmodelupdate = RuleAction{"skip_model_update"}
)

func (r RuleAction) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RuleAction) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RuleAction) String() string { _ = "STUB: not implemented"; return "" }
