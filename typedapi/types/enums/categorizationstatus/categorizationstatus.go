package categorizationstatus

type CategorizationStatus struct {
	Name string
}

var (
	Ok = CategorizationStatus{"ok"}

	Warn = CategorizationStatus{"warn"}
)

func (c CategorizationStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CategorizationStatus) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CategorizationStatus) String() string { _ = "STUB: not implemented"; return "" }
