package tdigestexecutionhint

type TDigestExecutionHint struct {
	Name string
}

var (
	Default = TDigestExecutionHint{"default"}

	Highaccuracy = TDigestExecutionHint{"high_accuracy"}
)

func (t TDigestExecutionHint) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TDigestExecutionHint) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TDigestExecutionHint) String() string { _ = "STUB: not implemented"; return "" }
