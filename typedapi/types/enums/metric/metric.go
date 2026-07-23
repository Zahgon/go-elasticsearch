package metric

type Metric struct {
	Name string
}

var (
	Min = Metric{"min"}

	Max = Metric{"max"}

	Sum = Metric{"sum"}

	Avg = Metric{"avg"}

	Valuecount = Metric{"value_count"}
)

func (m Metric) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Metric) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m Metric) String() string { _ = "STUB: not implemented"; return "" }
