package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pagerDutyEventProxy struct {
	v *types.PagerDutyEventProxy
}

func NewPagerDutyEventProxy() *_pagerDutyEventProxy { _ = "STUB: not implemented"; return nil }

func (s *_pagerDutyEventProxy) Host(host string) *_pagerDutyEventProxy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyEventProxy) Port(port int) *_pagerDutyEventProxy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyEventProxy) PagerDutyEventProxyCaster() *types.PagerDutyEventProxy {
	_ = "STUB: not implemented"
	return nil
}
