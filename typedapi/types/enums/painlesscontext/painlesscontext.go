package painlesscontext

type PainlessContext struct {
	Name string
}

var (
	Painlesstest = PainlessContext{"painless_test"}

	Filter = PainlessContext{"filter"}

	Score = PainlessContext{"score"}

	Booleanfield = PainlessContext{"boolean_field"}

	Datefield = PainlessContext{"date_field"}

	Doublefield = PainlessContext{"double_field"}

	Geopointfield = PainlessContext{"geo_point_field"}

	Ipfield = PainlessContext{"ip_field"}

	Keywordfield = PainlessContext{"keyword_field"}

	Longfield = PainlessContext{"long_field"}

	Compositefield = PainlessContext{"composite_field"}
)

func (p PainlessContext) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PainlessContext) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (p PainlessContext) String() string { _ = "STUB: not implemented"; return "" }
