package types

type CpuAcct struct {
	ControlGroup *string `json:"control_group,omitempty"`

	UsageNanos *int64 `json:"usage_nanos,omitempty"`
}

func (s *CpuAcct) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCpuAcct() *CpuAcct { _ = "STUB: not implemented"; return nil }
