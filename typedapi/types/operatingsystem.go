package types

type OperatingSystem struct {
	Cgroup    *Cgroup              `json:"cgroup,omitempty"`
	Cpu       *Cpu                 `json:"cpu,omitempty"`
	Mem       *ExtendedMemoryStats `json:"mem,omitempty"`
	Swap      *MemoryStats         `json:"swap,omitempty"`
	Timestamp *int64               `json:"timestamp,omitempty"`
}

func (s *OperatingSystem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewOperatingSystem() *OperatingSystem { _ = "STUB: not implemented"; return nil }
