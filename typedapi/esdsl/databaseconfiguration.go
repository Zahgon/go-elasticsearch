package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _databaseConfiguration struct {
	v *types.DatabaseConfiguration
}

func NewDatabaseConfiguration() *_databaseConfiguration { _ = "STUB: not implemented"; return nil }

func (s *_databaseConfiguration) Ipinfo(ipinfo types.IpinfoVariant) *_databaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_databaseConfiguration) Maxmind(maxmind types.MaxmindVariant) *_databaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_databaseConfiguration) Name(name string) *_databaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_databaseConfiguration) DatabaseConfigurationCaster() *types.DatabaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}
