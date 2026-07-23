package types

type NodeInfoOSCPU struct {
	CacheSize        string `json:"cache_size"`
	CacheSizeInBytes int    `json:"cache_size_in_bytes"`
	CoresPerSocket   int    `json:"cores_per_socket"`
	Mhz              int    `json:"mhz"`
	Model            string `json:"model"`
	TotalCores       int    `json:"total_cores"`
	TotalSockets     int    `json:"total_sockets"`
	Vendor           string `json:"vendor"`
}

func (s *NodeInfoOSCPU) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoOSCPU() *NodeInfoOSCPU { _ = "STUB: not implemented"; return nil }
