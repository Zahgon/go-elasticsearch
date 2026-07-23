package allocationexplaindecision

type AllocationExplainDecision struct {
	Name string
}

var (
	NO = AllocationExplainDecision{"NO"}

	YES = AllocationExplainDecision{"YES"}

	THROTTLE = AllocationExplainDecision{"THROTTLE"}

	ALWAYS = AllocationExplainDecision{"ALWAYS"}
)

func (a AllocationExplainDecision) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AllocationExplainDecision) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AllocationExplainDecision) String() string { _ = "STUB: not implemented"; return "" }
