package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/pagerdutyeventtype"
)

type _pagerDutyAction struct {
	v *types.PagerDutyAction
}

func NewPagerDutyAction(attachpayload bool, description string, incidentkey string) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) Account(account string) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) AttachPayload(attachpayload bool) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) Client(client string) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) ClientUrl(clienturl string) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) Contexts(contexts ...types.PagerDutyContextVariant) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) ContextsValues(contextsvalues []types.PagerDutyContext) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) Description(description string) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) EventType(eventtype pagerdutyeventtype.PagerDutyEventType) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) IncidentKey(incidentkey string) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) Proxy(proxy types.PagerDutyEventProxyVariant) *_pagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pagerDutyAction) PagerDutyActionCaster() *types.PagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}
