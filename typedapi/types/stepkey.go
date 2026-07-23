package types

type StepKey struct {
	Action *string `json:"action,omitempty"`

	Name  *string `json:"name,omitempty"`
	Phase string  `json:"phase"`
}

func (s *StepKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStepKey() *StepKey { _ = "STUB: not implemented"; return nil }

type StepKeyVariant interface {
	StepKeyCaster() *StepKey
}

func (s *StepKey) StepKeyCaster() *StepKey { _ = "STUB: not implemented"; return nil }
