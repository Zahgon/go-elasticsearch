package types

type ClusterJvmVersion struct {
	BundledJdk bool `json:"bundled_jdk"`

	Count int `json:"count"`

	UsingBundledJdk bool `json:"using_bundled_jdk"`

	Version string `json:"version"`

	VmName string `json:"vm_name"`

	VmVendor string `json:"vm_vendor"`

	VmVersion string `json:"vm_version"`
}

func (s *ClusterJvmVersion) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterJvmVersion() *ClusterJvmVersion { _ = "STUB: not implemented"; return nil }
