package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sourceOnlyRepository struct {
	v *types.SourceOnlyRepository
}

func NewSourceOnlyRepository() *_sourceOnlyRepository { _ = "STUB: not implemented"; return nil }

func (s *_sourceOnlyRepository) Settings(sourceonlyrepositorysettings types.SourceOnlyRepositorySettingsVariant) *_sourceOnlyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceOnlyRepository) Uuid(uuid string) *_sourceOnlyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceOnlyRepository) SourceOnlyRepositoryCaster() *types.SourceOnlyRepository {
	_ = "STUB: not implemented"
	return nil
}
