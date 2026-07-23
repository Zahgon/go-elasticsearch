package types

type FileSystemTotal struct {
	Available *string `json:"available,omitempty"`

	AvailableInBytes *int64 `json:"available_in_bytes,omitempty"`

	Free *string `json:"free,omitempty"`

	FreeInBytes *int64 `json:"free_in_bytes,omitempty"`

	Total *string `json:"total,omitempty"`

	TotalInBytes *int64 `json:"total_in_bytes,omitempty"`
}

func (s *FileSystemTotal) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFileSystemTotal() *FileSystemTotal { _ = "STUB: not implemented"; return nil }
