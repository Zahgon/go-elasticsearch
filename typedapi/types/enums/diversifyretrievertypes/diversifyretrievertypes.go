package diversifyretrievertypes

type DiversifyRetrieverTypes struct {
	Name string
}

var (
	Mmr = DiversifyRetrieverTypes{"mmr"}
)

func (d DiversifyRetrieverTypes) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DiversifyRetrieverTypes) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DiversifyRetrieverTypes) String() string { _ = "STUB: not implemented"; return "" }
