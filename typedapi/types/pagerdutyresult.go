package types

type PagerDutyResult struct {
	Event    PagerDutyEvent           `json:"event"`
	Reason   *string                  `json:"reason,omitempty"`
	Request  *HttpInputRequestResult  `json:"request,omitempty"`
	Response *HttpInputResponseResult `json:"response,omitempty"`
}

func (s *PagerDutyResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPagerDutyResult() *PagerDutyResult { _ = "STUB: not implemented"; return nil }
