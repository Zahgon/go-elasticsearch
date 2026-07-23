package cohereinputtype

type CohereInputType struct {
	Name string
}

var (
	Classification = CohereInputType{"classification"}

	Clustering = CohereInputType{"clustering"}

	Ingest = CohereInputType{"ingest"}

	Search = CohereInputType{"search"}
)

func (c CohereInputType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CohereInputType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CohereInputType) String() string { _ = "STUB: not implemented"; return "" }
