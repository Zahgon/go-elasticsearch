package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/pagerdutyeventtype"
)

type PagerDutyAction struct {
	Account       *string                                `json:"account,omitempty"`
	AttachPayload bool                                   `json:"attach_payload"`
	Client        *string                                `json:"client,omitempty"`
	ClientUrl     *string                                `json:"client_url,omitempty"`
	Contexts      []PagerDutyContext                     `json:"contexts,omitempty"`
	Description   string                                 `json:"description"`
	EventType     *pagerdutyeventtype.PagerDutyEventType `json:"event_type,omitempty"`
	IncidentKey   string                                 `json:"incident_key"`
	Proxy         *PagerDutyEventProxy                   `json:"proxy,omitempty"`
}

func (s *PagerDutyAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPagerDutyAction() *PagerDutyAction { _ = "STUB: not implemented"; return nil }

type PagerDutyActionVariant interface {
	PagerDutyActionCaster() *PagerDutyAction
}

func (s *PagerDutyAction) PagerDutyActionCaster() *PagerDutyAction {
	_ = "STUB: not implemented"
	return nil
}
