package useragentproperty

type UserAgentProperty struct {
	Name string
}

var (
	Name = UserAgentProperty{"name"}

	Os = UserAgentProperty{"os"}

	Device = UserAgentProperty{"device"}

	Original = UserAgentProperty{"original"}

	Version = UserAgentProperty{"version"}
)

func (u UserAgentProperty) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *UserAgentProperty) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (u UserAgentProperty) String() string { _ = "STUB: not implemented"; return "" }
