package types

type FileCountSnapshotStats struct {
	FileCount   int   `json:"file_count"`
	SizeInBytes int64 `json:"size_in_bytes"`
}

func (s *FileCountSnapshotStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFileCountSnapshotStats() *FileCountSnapshotStats { _ = "STUB: not implemented"; return nil }
