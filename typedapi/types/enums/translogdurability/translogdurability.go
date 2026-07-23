package translogdurability

type TranslogDurability struct {
	Name string
}

var (
	Request = TranslogDurability{"request"}

	Async = TranslogDurability{"async"}
)

func (t TranslogDurability) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TranslogDurability) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TranslogDurability) String() string { _ = "STUB: not implemented"; return "" }
