package types

type ClusterOperatingSystemArchitecture struct {
	Arch string `json:"arch"`

	Count int `json:"count"`
}

func (s *ClusterOperatingSystemArchitecture) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterOperatingSystemArchitecture() *ClusterOperatingSystemArchitecture {
	_ = "STUB: not implemented"
	return nil
}
