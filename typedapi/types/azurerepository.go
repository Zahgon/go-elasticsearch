package types

type AzureRepository struct {
	Settings *AzureRepositorySettings `json:"settings,omitempty"`

	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *AzureRepository) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s AzureRepository) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewAzureRepository() *AzureRepository { _ = "STUB: not implemented"; return nil }

type AzureRepositoryVariant interface {
	AzureRepositoryCaster() *AzureRepository
}

func (s *AzureRepository) AzureRepositoryCaster() *AzureRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *AzureRepository) RepositoryCaster() *Repository { _ = "STUB: not implemented"; return nil }
