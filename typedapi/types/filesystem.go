package types

type FileSystem struct {
	Data []DataPathStats `json:"data,omitempty"`

	IoStats *IoStats `json:"io_stats,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`

	Total *FileSystemTotal `json:"total,omitempty"`
}

func (s *FileSystem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFileSystem() *FileSystem { _ = "STUB: not implemented"; return nil }
