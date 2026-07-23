package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ipinfo struct {
	v *types.Ipinfo
}

func NewIpinfo() *_ipinfo { _ = "STUB: not implemented"; return nil }

func (s *_ipinfo) DatabaseConfigurationCaster() *types.DatabaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipinfo) DatabaseConfigurationFullCaster() *types.DatabaseConfigurationFull {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipinfo) IpinfoCaster() *types.Ipinfo { _ = "STUB: not implemented"; return nil }
