package types

type HdrMethod struct {
	NumberOfSignificantValueDigits *int `json:"number_of_significant_value_digits,omitempty"`
}

func (s *HdrMethod) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHdrMethod() *HdrMethod { _ = "STUB: not implemented"; return nil }

type HdrMethodVariant interface {
	HdrMethodCaster() *HdrMethod
}

func (s *HdrMethod) HdrMethodCaster() *HdrMethod { _ = "STUB: not implemented"; return nil }
