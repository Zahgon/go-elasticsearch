package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _gcsRepository struct {
	v *types.GcsRepository
}

func NewGcsRepository(settings types.GcsRepositorySettingsVariant) *_gcsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepository) Settings(settings types.GcsRepositorySettingsVariant) *_gcsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepository) Uuid(uuid string) *_gcsRepository { _ = "STUB: not implemented"; return nil }

func (s *_gcsRepository) GcsRepositoryCaster() *types.GcsRepository {
	_ = "STUB: not implemented"
	return nil
}
