package licensestatus

type LicenseStatus struct {
	Name string
}

var (
	Active = LicenseStatus{"active"}

	Valid = LicenseStatus{"valid"}

	Invalid = LicenseStatus{"invalid"}

	Expired = LicenseStatus{"expired"}
)

func (l LicenseStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LicenseStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l LicenseStatus) String() string { _ = "STUB: not implemented"; return "" }
