package types

type GcsRepository struct {
	Settings GcsRepositorySettings `json:"settings"`

	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *GcsRepository) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GcsRepository) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGcsRepository() *GcsRepository { _ = "STUB: not implemented"; return nil }

type GcsRepositoryVariant interface {
	GcsRepositoryCaster() *GcsRepository
}

func (s *GcsRepository) GcsRepositoryCaster() *GcsRepository { _ = "STUB: not implemented"; return nil }

func (s *GcsRepository) RepositoryCaster() *Repository { _ = "STUB: not implemented"; return nil }
