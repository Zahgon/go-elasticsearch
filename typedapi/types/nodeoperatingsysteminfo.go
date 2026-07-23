package types

type NodeOperatingSystemInfo struct {
	AllocatedProcessors *int `json:"allocated_processors,omitempty"`

	Arch string `json:"arch"`

	AvailableProcessors int             `json:"available_processors"`
	Cpu                 *NodeInfoOSCPU  `json:"cpu,omitempty"`
	Mem                 *NodeInfoMemory `json:"mem,omitempty"`

	Name       string `json:"name"`
	PrettyName string `json:"pretty_name"`

	RefreshIntervalInMillis int64           `json:"refresh_interval_in_millis"`
	Swap                    *NodeInfoMemory `json:"swap,omitempty"`

	Version string `json:"version"`
}

func (s *NodeOperatingSystemInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeOperatingSystemInfo() *NodeOperatingSystemInfo { _ = "STUB: not implemented"; return nil }
