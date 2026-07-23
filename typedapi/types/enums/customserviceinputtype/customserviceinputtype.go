package customserviceinputtype

type CustomServiceInputType struct {
	Name string
}

var (
	Classification = CustomServiceInputType{"classification"}

	Clustering = CustomServiceInputType{"clustering"}

	Ingest = CustomServiceInputType{"ingest"}

	Search = CustomServiceInputType{"search"}
)

func (c CustomServiceInputType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CustomServiceInputType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CustomServiceInputType) String() string { _ = "STUB: not implemented"; return "" }
