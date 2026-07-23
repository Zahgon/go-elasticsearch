package types

type PagerDutyEventProxy struct {
	Host *string `json:"host,omitempty"`
	Port *int    `json:"port,omitempty"`
}

func (s *PagerDutyEventProxy) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPagerDutyEventProxy() *PagerDutyEventProxy { _ = "STUB: not implemented"; return nil }

type PagerDutyEventProxyVariant interface {
	PagerDutyEventProxyCaster() *PagerDutyEventProxy
}

func (s *PagerDutyEventProxy) PagerDutyEventProxyCaster() *PagerDutyEventProxy {
	_ = "STUB: not implemented"
	return nil
}
