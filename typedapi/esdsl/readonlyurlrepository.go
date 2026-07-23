package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _readOnlyUrlRepository struct {
	v *types.ReadOnlyUrlRepository
}

func NewReadOnlyUrlRepository(settings types.ReadOnlyUrlRepositorySettingsVariant) *_readOnlyUrlRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepository) Settings(settings types.ReadOnlyUrlRepositorySettingsVariant) *_readOnlyUrlRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepository) Uuid(uuid string) *_readOnlyUrlRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepository) ReadOnlyUrlRepositoryCaster() *types.ReadOnlyUrlRepository {
	_ = "STUB: not implemented"
	return nil
}
