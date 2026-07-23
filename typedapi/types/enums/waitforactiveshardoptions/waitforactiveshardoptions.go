package waitforactiveshardoptions

type WaitForActiveShardOptions struct {
	Name string
}

var (
	All = WaitForActiveShardOptions{"all"}

	IndexSetting = WaitForActiveShardOptions{"index-setting"}
)

func (w WaitForActiveShardOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WaitForActiveShardOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (w WaitForActiveShardOptions) String() string { _ = "STUB: not implemented"; return "" }
