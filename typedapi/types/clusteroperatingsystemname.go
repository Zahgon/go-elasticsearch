package types

type ClusterOperatingSystemName struct {
	Count int `json:"count"`

	Name string `json:"name"`
}

func (s *ClusterOperatingSystemName) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterOperatingSystemName() *ClusterOperatingSystemName {
	_ = "STUB: not implemented"
	return nil
}
