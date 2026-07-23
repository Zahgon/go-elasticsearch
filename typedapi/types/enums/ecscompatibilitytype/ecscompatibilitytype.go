package ecscompatibilitytype

type EcsCompatibilityType struct {
	Name string
}

var (
	Disabled = EcsCompatibilityType{"disabled"}

	V1 = EcsCompatibilityType{"v1"}
)

func (e EcsCompatibilityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EcsCompatibilityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e EcsCompatibilityType) String() string { _ = "STUB: not implemented"; return "" }
