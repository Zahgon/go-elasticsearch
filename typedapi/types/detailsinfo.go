package types

type DetailsInfo struct {
	Blob BlobDetails `json:"blob"`

	OverwriteElapsed Duration `json:"overwrite_elapsed,omitempty"`

	OverwriteElapsedNanos *int64 `json:"overwrite_elapsed_nanos,omitempty"`

	WriteElapsed Duration `json:"write_elapsed"`

	WriteElapsedNanos int64 `json:"write_elapsed_nanos"`

	WriteThrottled Duration `json:"write_throttled"`

	WriteThrottledNanos int64 `json:"write_throttled_nanos"`

	WriterNode SnapshotNodeInfo `json:"writer_node"`
}

func (s *DetailsInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDetailsInfo() *DetailsInfo { _ = "STUB: not implemented"; return nil }
