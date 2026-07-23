package sqlformat

type SqlFormat struct {
	Name string
}

var (
	Csv = SqlFormat{"csv"}

	Json = SqlFormat{"json"}

	Tsv = SqlFormat{"tsv"}

	Txt = SqlFormat{"txt"}

	Yaml = SqlFormat{"yaml"}

	Cbor = SqlFormat{"cbor"}

	Smile = SqlFormat{"smile"}
)

func (s SqlFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SqlFormat) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SqlFormat) String() string { _ = "STUB: not implemented"; return "" }
