package types

type ClusterProcess struct {
	Cpu ClusterProcessCpu `json:"cpu"`

	OpenFileDescriptors ClusterProcessOpenFileDescriptors `json:"open_file_descriptors"`
}

func NewClusterProcess() *ClusterProcess { _ = "STUB: not implemented"; return nil }
