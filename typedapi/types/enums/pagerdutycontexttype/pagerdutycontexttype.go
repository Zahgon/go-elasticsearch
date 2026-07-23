package pagerdutycontexttype

type PagerDutyContextType struct {
	Name string
}

var (
	Link = PagerDutyContextType{"link"}

	Image = PagerDutyContextType{"image"}
)

func (p PagerDutyContextType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PagerDutyContextType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PagerDutyContextType) String() string { _ = "STUB: not implemented"; return "" }
