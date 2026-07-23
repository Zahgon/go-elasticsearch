package jsonprocessorconflictstrategy

type JsonProcessorConflictStrategy struct {
	Name string
}

var (
	Replace = JsonProcessorConflictStrategy{"replace"}

	Merge = JsonProcessorConflictStrategy{"merge"}
)

func (j JsonProcessorConflictStrategy) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JsonProcessorConflictStrategy) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JsonProcessorConflictStrategy) String() string { _ = "STUB: not implemented"; return "" }
