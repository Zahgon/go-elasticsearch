package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _repository struct {
	v types.Repository
}

func NewRepository() *_repository { _ = "STUB: not implemented"; return nil }

func (u *_repository) UnknownRepository(unknown json.RawMessage) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_repository) AzureRepository(azurerepository types.AzureRepositoryVariant) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_azureRepository) RepositoryCaster() *types.Repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_repository) GcsRepository(gcsrepository types.GcsRepositoryVariant) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_gcsRepository) RepositoryCaster() *types.Repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_repository) S3Repository(s3repository types.S3RepositoryVariant) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_s3Repository) RepositoryCaster() *types.Repository { _ = "STUB: not implemented"; return nil }

func (u *_repository) SharedFileSystemRepository(sharedfilesystemrepository types.SharedFileSystemRepositoryVariant) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sharedFileSystemRepository) RepositoryCaster() *types.Repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_repository) ReadOnlyUrlRepository(readonlyurlrepository types.ReadOnlyUrlRepositoryVariant) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_readOnlyUrlRepository) RepositoryCaster() *types.Repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_repository) SourceOnlyRepository(sourceonlyrepository types.SourceOnlyRepositoryVariant) *_repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sourceOnlyRepository) RepositoryCaster() *types.Repository {
	_ = "STUB: not implemented"
	return nil
}

func (u *_repository) RepositoryCaster() *types.Repository { _ = "STUB: not implemented"; return nil }
