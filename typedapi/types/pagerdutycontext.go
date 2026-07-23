package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/pagerdutycontexttype"
)

type PagerDutyContext struct {
	Href *string                                   `json:"href,omitempty"`
	Src  *string                                   `json:"src,omitempty"`
	Type pagerdutycontexttype.PagerDutyContextType `json:"type"`
}

func (s *PagerDutyContext) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPagerDutyContext() *PagerDutyContext { _ = "STUB: not implemented"; return nil }

type PagerDutyContextVariant interface {
	PagerDutyContextCaster() *PagerDutyContext
}

func (s *PagerDutyContext) PagerDutyContextCaster() *PagerDutyContext {
	_ = "STUB: not implemented"
	return nil
}
