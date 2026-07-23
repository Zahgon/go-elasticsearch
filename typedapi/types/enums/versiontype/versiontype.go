package versiontype

type VersionType struct {
	Name string
}

var (
	Internal = VersionType{"internal"}

	External = VersionType{"external"}

	Externalgte = VersionType{"external_gte"}
)

func (v VersionType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VersionType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (v VersionType) String() string { _ = "STUB: not implemented"; return "" }
