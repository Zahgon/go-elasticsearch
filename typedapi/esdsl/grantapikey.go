package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _grantApiKey struct {
	v *types.GrantApiKey
}

func NewGrantApiKey() *_grantApiKey { _ = "STUB: not implemented"; return nil }

func (s *_grantApiKey) Expiration(durationlarge string) *_grantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grantApiKey) Metadata(metadata types.MetadataVariant) *_grantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grantApiKey) Name(name string) *_grantApiKey { _ = "STUB: not implemented"; return nil }

func (s *_grantApiKey) RoleDescriptors(roledescriptors []map[string]types.RoleDescriptor) *_grantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grantApiKey) GrantApiKeyCaster() *types.GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}
