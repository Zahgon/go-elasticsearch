package findstructureformat

type FindStructureFormat struct {
	Name string
}

var (
	Ndjson = FindStructureFormat{"ndjson"}

	Xml = FindStructureFormat{"xml"}

	Delimited = FindStructureFormat{"delimited"}

	Semistructuredtext = FindStructureFormat{"semi_structured_text"}
)

func (f FindStructureFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FindStructureFormat) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FindStructureFormat) String() string { _ = "STUB: not implemented"; return "" }
