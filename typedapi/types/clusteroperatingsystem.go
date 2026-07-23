package types

type ClusterOperatingSystem struct {
	AllocatedProcessors int `json:"allocated_processors"`

	Architectures []ClusterOperatingSystemArchitecture `json:"architectures,omitempty"`

	AvailableProcessors int `json:"available_processors"`

	Mem OperatingSystemMemoryInfo `json:"mem"`

	Names []ClusterOperatingSystemName `json:"names"`

	PrettyNames []ClusterOperatingSystemPrettyName `json:"pretty_names"`
}

func (s *ClusterOperatingSystem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterOperatingSystem() *ClusterOperatingSystem { _ = "STUB: not implemented"; return nil }
