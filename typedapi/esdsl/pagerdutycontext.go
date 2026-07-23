package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/pagerdutycontexttype"
)

type _pagerDutyContext struct {
	v *types.PagerDutyContext
}

func NewPagerDutyContext(type_ pagerdutycontexttype.PagerDutyContextType) *_pagerDutyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyContext) Href(href string) *_pagerDutyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyContext) Src(src string) *_pagerDutyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyContext) Type(type_ pagerdutycontexttype.PagerDutyContextType) *_pagerDutyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyContext) PagerDutyContextCaster() *types.PagerDutyContext {
	_ = "STUB: not implemented"
	return nil
}
