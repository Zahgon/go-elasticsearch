package sorttype

type SortType struct {
	Name string
}

var (
	Doc = SortType{"_doc"}

	Geodistance = SortType{"_geo_distance"}

	Score = SortType{"_score"}

	Script = SortType{"_script"}

	Fieldsort = SortType{"field_sort"}
)

func (s SortType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SortType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SortType) String() string { _ = "STUB: not implemented"; return "" }
