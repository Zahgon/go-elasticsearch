package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _maxmind struct {
	v *types.Maxmind
}

func NewMaxmind() *_maxmind { _ = "STUB: not implemented"; return nil }

func (s *_maxmind) AccountId(id string) *_maxmind { _ = "STUB: not implemented"; return nil }

func (s *_maxmind) DatabaseConfigurationCaster() *types.DatabaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxmind) DatabaseConfigurationFullCaster() *types.DatabaseConfigurationFull {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxmind) MaxmindCaster() *types.Maxmind { _ = "STUB: not implemented"; return nil }
