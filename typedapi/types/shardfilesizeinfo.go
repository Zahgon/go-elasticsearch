package types

type ShardFileSizeInfo struct {
	AverageSizeInBytes *int64 `json:"average_size_in_bytes,omitempty"`
	Count              *int64 `json:"count,omitempty"`
	Description        string `json:"description"`
	MaxSizeInBytes     *int64 `json:"max_size_in_bytes,omitempty"`
	MinSizeInBytes     *int64 `json:"min_size_in_bytes,omitempty"`
	SizeInBytes        int64  `json:"size_in_bytes"`
}

func (s *ShardFileSizeInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardFileSizeInfo() *ShardFileSizeInfo { _ = "STUB: not implemented"; return nil }
