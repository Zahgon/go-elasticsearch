package samplingmethod

type SamplingMethod struct {
	Name string
}

var (
	Aggregate = SamplingMethod{"aggregate"}

	Lastvalue = SamplingMethod{"last_value"}
)

func (s SamplingMethod) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SamplingMethod) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SamplingMethod) String() string { _ = "STUB: not implemented"; return "" }
