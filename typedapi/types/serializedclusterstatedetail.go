package types

type SerializedClusterStateDetail struct {
	CompressedSize          *string `json:"compressed_size,omitempty"`
	CompressedSizeInBytes   *int64  `json:"compressed_size_in_bytes,omitempty"`
	Count                   *int64  `json:"count,omitempty"`
	UncompressedSize        *string `json:"uncompressed_size,omitempty"`
	UncompressedSizeInBytes *int64  `json:"uncompressed_size_in_bytes,omitempty"`
}

func (s *SerializedClusterStateDetail) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSerializedClusterStateDetail() *SerializedClusterStateDetail {
	_ = "STUB: not implemented"
	return nil
}
