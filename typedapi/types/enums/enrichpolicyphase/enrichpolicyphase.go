package enrichpolicyphase

type EnrichPolicyPhase struct {
	Name string
}

var (
	SCHEDULED = EnrichPolicyPhase{"SCHEDULED"}

	RUNNING = EnrichPolicyPhase{"RUNNING"}

	COMPLETE = EnrichPolicyPhase{"COMPLETE"}

	FAILED = EnrichPolicyPhase{"FAILED"}

	CANCELLED = EnrichPolicyPhase{"CANCELLED"}
)

func (e EnrichPolicyPhase) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EnrichPolicyPhase) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e EnrichPolicyPhase) String() string { _ = "STUB: not implemented"; return "" }
