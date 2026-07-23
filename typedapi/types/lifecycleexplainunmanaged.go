package types

type LifecycleExplainUnmanaged struct {
	Index   string `json:"index"`
	Managed bool   `json:"managed,omitempty"`
}

func (s *LifecycleExplainUnmanaged) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LifecycleExplainUnmanaged) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLifecycleExplainUnmanaged() *LifecycleExplainUnmanaged {
	_ = "STUB: not implemented"
	return nil
}
