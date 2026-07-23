package types

type CleanupRepositoryResults struct {
	DeletedBlobs int64 `json:"deleted_blobs"`

	DeletedBytes int64 `json:"deleted_bytes"`
}

func (s *CleanupRepositoryResults) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCleanupRepositoryResults() *CleanupRepositoryResults { _ = "STUB: not implemented"; return nil }
