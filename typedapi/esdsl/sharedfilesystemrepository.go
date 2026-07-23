package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sharedFileSystemRepository struct {
	v *types.SharedFileSystemRepository
}

func NewSharedFileSystemRepository(settings types.SharedFileSystemRepositorySettingsVariant) *_sharedFileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepository) Settings(settings types.SharedFileSystemRepositorySettingsVariant) *_sharedFileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepository) Uuid(uuid string) *_sharedFileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepository) SharedFileSystemRepositoryCaster() *types.SharedFileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}
