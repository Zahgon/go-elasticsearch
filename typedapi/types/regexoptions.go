package types

type RegexOptions struct {
	Flags *string `json:"flags,omitempty"`

	MaxDeterminizedStates *int `json:"max_determinized_states,omitempty"`
}

func (s *RegexOptions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRegexOptions() *RegexOptions { _ = "STUB: not implemented"; return nil }

type RegexOptionsVariant interface {
	RegexOptionsCaster() *RegexOptions
}

func (s *RegexOptions) RegexOptionsCaster() *RegexOptions { _ = "STUB: not implemented"; return nil }
