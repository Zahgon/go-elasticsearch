package esqlformat

type EsqlFormat struct {
	Name string
}

var (
	Csv = EsqlFormat{"csv"}

	Json = EsqlFormat{"json"}

	Tsv = EsqlFormat{"tsv"}

	Txt = EsqlFormat{"txt"}

	Yaml = EsqlFormat{"yaml"}

	Cbor = EsqlFormat{"cbor"}

	Smile = EsqlFormat{"smile"}

	Arrow = EsqlFormat{"arrow"}
)

func (e EsqlFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EsqlFormat) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e EsqlFormat) String() string { _ = "STUB: not implemented"; return "" }
