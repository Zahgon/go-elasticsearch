package types

type Process struct {
	Cpu *Cpu `json:"cpu,omitempty"`

	MaxFileDescriptors *int `json:"max_file_descriptors,omitempty"`

	Mem *MemoryStats `json:"mem,omitempty"`

	OpenFileDescriptors *int `json:"open_file_descriptors,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (s *Process) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewProcess() *Process { _ = "STUB: not implemented"; return nil }
