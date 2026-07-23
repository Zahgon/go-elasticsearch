package types

type ClusterOperatingSystemPrettyName struct {
	Count int `json:"count"`

	PrettyName string `json:"pretty_name"`
}

func (s *ClusterOperatingSystemPrettyName) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterOperatingSystemPrettyName() *ClusterOperatingSystemPrettyName {
	_ = "STUB: not implemented"
	return nil
}
