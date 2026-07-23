package types

type CgroupMemory struct {
	ControlGroup *string `json:"control_group,omitempty"`

	LimitInBytes *string `json:"limit_in_bytes,omitempty"`

	UsageInBytes *string `json:"usage_in_bytes,omitempty"`
}

func (s *CgroupMemory) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCgroupMemory() *CgroupMemory { _ = "STUB: not implemented"; return nil }
