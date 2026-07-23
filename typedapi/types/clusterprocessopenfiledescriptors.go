package types

type ClusterProcessOpenFileDescriptors struct {
	Avg int64 `json:"avg"`

	Max int64 `json:"max"`

	Min int64 `json:"min"`
}

func (s *ClusterProcessOpenFileDescriptors) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterProcessOpenFileDescriptors() *ClusterProcessOpenFileDescriptors {
	_ = "STUB: not implemented"
	return nil
}
