package types

type ClusterJvm struct {
	MaxUptime Duration `json:"max_uptime,omitempty"`

	MaxUptimeInMillis int64 `json:"max_uptime_in_millis"`

	Mem ClusterJvmMemory `json:"mem"`

	Threads int64 `json:"threads"`

	Versions []ClusterJvmVersion `json:"versions"`
}

func (s *ClusterJvm) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterJvm() *ClusterJvm { _ = "STUB: not implemented"; return nil }
