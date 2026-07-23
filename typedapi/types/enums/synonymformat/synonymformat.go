package synonymformat

type SynonymFormat struct {
	Name string
}

var (
	Solr = SynonymFormat{"solr"}

	Wordnet = SynonymFormat{"wordnet"}
)

func (s SynonymFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SynonymFormat) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SynonymFormat) String() string { _ = "STUB: not implemented"; return "" }
