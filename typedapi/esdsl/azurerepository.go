package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureRepository struct {
	v *types.AzureRepository
}

func NewAzureRepository() *_azureRepository { _ = "STUB: not implemented"; return nil }

func (s *_azureRepository) Settings(settings types.AzureRepositorySettingsVariant) *_azureRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepository) Uuid(uuid string) *_azureRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepository) AzureRepositoryCaster() *types.AzureRepository {
	_ = "STUB: not implemented"
	return nil
}
