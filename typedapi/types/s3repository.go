package types

type S3Repository struct {
	Settings S3RepositorySettings `json:"settings"`

	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *S3Repository) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s S3Repository) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewS3Repository() *S3Repository { _ = "STUB: not implemented"; return nil }

type S3RepositoryVariant interface {
	S3RepositoryCaster() *S3Repository
}

func (s *S3Repository) S3RepositoryCaster() *S3Repository { _ = "STUB: not implemented"; return nil }

func (s *S3Repository) RepositoryCaster() *Repository { _ = "STUB: not implemented"; return nil }
