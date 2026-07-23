package stringdistance

type StringDistance struct {
	Name string
}

var (
	Internal = StringDistance{"internal"}

	Dameraulevenshtein = StringDistance{"damerau_levenshtein"}

	Levenshtein = StringDistance{"levenshtein"}

	Jarowinkler = StringDistance{"jaro_winkler"}

	Ngram = StringDistance{"ngram"}
)

func (s StringDistance) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StringDistance) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s StringDistance) String() string { _ = "STUB: not implemented"; return "" }
