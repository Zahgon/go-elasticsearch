package simplequerystringflag

type SimpleQueryStringFlag struct {
	Name string
}

var (
	NONE = SimpleQueryStringFlag{"NONE"}

	AND = SimpleQueryStringFlag{"AND"}

	NOT = SimpleQueryStringFlag{"NOT"}

	OR = SimpleQueryStringFlag{"OR"}

	PREFIX = SimpleQueryStringFlag{"PREFIX"}

	PHRASE = SimpleQueryStringFlag{"PHRASE"}

	PRECEDENCE = SimpleQueryStringFlag{"PRECEDENCE"}

	ESCAPE = SimpleQueryStringFlag{"ESCAPE"}

	WHITESPACE = SimpleQueryStringFlag{"WHITESPACE"}

	FUZZY = SimpleQueryStringFlag{"FUZZY"}

	NEAR = SimpleQueryStringFlag{"NEAR"}

	SLOP = SimpleQueryStringFlag{"SLOP"}

	ALL = SimpleQueryStringFlag{"ALL"}
)

func (s SimpleQueryStringFlag) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SimpleQueryStringFlag) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SimpleQueryStringFlag) String() string { _ = "STUB: not implemented"; return "" }
