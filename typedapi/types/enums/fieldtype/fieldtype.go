package fieldtype

type FieldType struct {
	Name string
}

var (
	None = FieldType{"none"}

	Geopoint = FieldType{"geo_point"}

	Geoshape = FieldType{"geo_shape"}

	Ip = FieldType{"ip"}

	Binary = FieldType{"binary"}

	Keyword = FieldType{"keyword"}

	Text = FieldType{"text"}

	Searchasyoutype = FieldType{"search_as_you_type"}

	Wildcard = FieldType{"wildcard"}

	Date = FieldType{"date"}

	Datenanos = FieldType{"date_nanos"}

	Boolean = FieldType{"boolean"}

	Completion = FieldType{"completion"}

	Nested = FieldType{"nested"}

	Object = FieldType{"object"}

	Passthrough = FieldType{"passthrough"}

	Version = FieldType{"version"}

	Murmur3 = FieldType{"murmur3"}

	Tokencount = FieldType{"token_count"}

	Percolator = FieldType{"percolator"}

	Integer = FieldType{"integer"}

	Long = FieldType{"long"}

	Short = FieldType{"short"}

	Byte = FieldType{"byte"}

	Float = FieldType{"float"}

	Halffloat = FieldType{"half_float"}

	Scaledfloat = FieldType{"scaled_float"}

	Double = FieldType{"double"}

	Integerrange = FieldType{"integer_range"}

	Floatrange = FieldType{"float_range"}

	Longrange = FieldType{"long_range"}

	Doublerange = FieldType{"double_range"}

	Daterange = FieldType{"date_range"}

	Iprange = FieldType{"ip_range"}

	Alias = FieldType{"alias"}

	Join = FieldType{"join"}

	Rankfeature = FieldType{"rank_feature"}

	Rankfeatures = FieldType{"rank_features"}

	Flattened = FieldType{"flattened"}

	Shape = FieldType{"shape"}

	Histogram = FieldType{"histogram"}

	Constantkeyword = FieldType{"constant_keyword"}

	Countedkeyword = FieldType{"counted_keyword"}

	Aggregatemetricdouble = FieldType{"aggregate_metric_double"}

	Densevector = FieldType{"dense_vector"}

	Semantictext = FieldType{"semantic_text"}

	Sparsevector = FieldType{"sparse_vector"}

	Matchonlytext = FieldType{"match_only_text"}

	Icucollationkeyword = FieldType{"icu_collation_keyword"}
)

func (f FieldType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FieldType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FieldType) String() string { _ = "STUB: not implemented"; return "" }
