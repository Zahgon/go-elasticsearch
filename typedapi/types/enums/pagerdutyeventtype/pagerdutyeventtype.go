package pagerdutyeventtype

type PagerDutyEventType struct {
	Name string
}

var (
	Trigger = PagerDutyEventType{"trigger"}

	Resolve = PagerDutyEventType{"resolve"}

	Acknowledge = PagerDutyEventType{"acknowledge"}
)

func (p PagerDutyEventType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PagerDutyEventType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PagerDutyEventType) String() string { _ = "STUB: not implemented"; return "" }
