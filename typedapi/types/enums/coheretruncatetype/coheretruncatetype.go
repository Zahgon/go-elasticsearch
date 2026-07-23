package coheretruncatetype

type CohereTruncateType struct {
	Name string
}

var (
	END = CohereTruncateType{"END"}

	NONE = CohereTruncateType{"NONE"}

	START = CohereTruncateType{"START"}
)

func (c CohereTruncateType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CohereTruncateType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CohereTruncateType) String() string { _ = "STUB: not implemented"; return "" }
