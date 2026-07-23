package failurestorestatus

type FailureStoreStatus struct {
	Name string
}

var (
	Notapplicableorunknown = FailureStoreStatus{"not_applicable_or_unknown"}

	Used = FailureStoreStatus{"used"}

	Notenabled = FailureStoreStatus{"not_enabled"}

	Failed = FailureStoreStatus{"failed"}
)

func (f FailureStoreStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FailureStoreStatus) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FailureStoreStatus) String() string { _ = "STUB: not implemented"; return "" }
