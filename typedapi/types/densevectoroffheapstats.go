package types

type DenseVectorOffHeapStats struct {
	Fielddata           map[string]map[string]int64 `json:"fielddata,omitempty"`
	TotalCenifSize      ByteSize                    `json:"total_cenif_size,omitempty"`
	TotalCenifSizeBytes int64                       `json:"total_cenif_size_bytes"`
	TotalClivfSize      ByteSize                    `json:"total_clivf_size,omitempty"`
	TotalClivfSizeBytes int64                       `json:"total_clivf_size_bytes"`
	TotalSize           ByteSize                    `json:"total_size,omitempty"`
	TotalSizeBytes      int64                       `json:"total_size_bytes"`
	TotalVebSize        ByteSize                    `json:"total_veb_size,omitempty"`
	TotalVebSizeBytes   int64                       `json:"total_veb_size_bytes"`
	TotalVecSize        ByteSize                    `json:"total_vec_size,omitempty"`
	TotalVecSizeBytes   int64                       `json:"total_vec_size_bytes"`
	TotalVeqSize        ByteSize                    `json:"total_veq_size,omitempty"`
	TotalVeqSizeBytes   int64                       `json:"total_veq_size_bytes"`
	TotalVexSize        ByteSize                    `json:"total_vex_size,omitempty"`
	TotalVexSizeBytes   int64                       `json:"total_vex_size_bytes"`
}

func (s *DenseVectorOffHeapStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDenseVectorOffHeapStats() *DenseVectorOffHeapStats { _ = "STUB: not implemented"; return nil }
