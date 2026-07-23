package policytype

type PolicyType struct {
	Name string
}

var (
	Geomatch = PolicyType{"geo_match"}

	Match = PolicyType{"match"}

	Range = PolicyType{"range"}
)

func (p PolicyType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PolicyType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (p PolicyType) String() string { _ = "STUB: not implemented"; return "" }
