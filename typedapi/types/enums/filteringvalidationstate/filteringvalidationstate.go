package filteringvalidationstate

type FilteringValidationState struct {
	Name string
}

var (
	Edited = FilteringValidationState{"edited"}

	Invalid = FilteringValidationState{"invalid"}

	Valid = FilteringValidationState{"valid"}
)

func (f FilteringValidationState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilteringValidationState) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FilteringValidationState) String() string { _ = "STUB: not implemented"; return "" }
