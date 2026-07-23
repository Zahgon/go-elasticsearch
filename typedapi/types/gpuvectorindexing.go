package types

type GpuVectorIndexing struct {
	Available bool `json:"available"`
	Enabled   bool `json:"enabled"`

	IndexBuildCount int64 `json:"index_build_count"`

	Nodes []GpuNodeStats `json:"nodes"`

	NodesWithGpu int `json:"nodes_with_gpu"`
}

func (s *GpuVectorIndexing) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGpuVectorIndexing() *GpuVectorIndexing { _ = "STUB: not implemented"; return nil }
