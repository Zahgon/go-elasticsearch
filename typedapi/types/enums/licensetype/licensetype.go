package licensetype

type LicenseType struct {
	Name string
}

var (
	Missing = LicenseType{"missing"}

	Trial = LicenseType{"trial"}

	Basic = LicenseType{"basic"}

	Standard = LicenseType{"standard"}

	Dev = LicenseType{"dev"}

	Silver = LicenseType{"silver"}

	Gold = LicenseType{"gold"}

	Platinum = LicenseType{"platinum"}

	Enterprise = LicenseType{"enterprise"}
)

func (l LicenseType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LicenseType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l LicenseType) String() string { _ = "STUB: not implemented"; return "" }
