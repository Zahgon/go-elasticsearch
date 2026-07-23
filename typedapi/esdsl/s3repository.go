package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _s3Repository struct {
	v *types.S3Repository
}

func NewS3Repository(settings types.S3RepositorySettingsVariant) *_s3Repository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3Repository) Settings(settings types.S3RepositorySettingsVariant) *_s3Repository {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3Repository) Uuid(uuid string) *_s3Repository { _ = "STUB: not implemented"; return nil }

func (s *_s3Repository) S3RepositoryCaster() *types.S3Repository {
	_ = "STUB: not implemented"
	return nil
}
