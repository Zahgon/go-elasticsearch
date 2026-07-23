package types

type Calendar struct {
	CalendarId string `json:"calendar_id"`

	Description *string `json:"description,omitempty"`

	JobIds []string `json:"job_ids"`
}

func (s *Calendar) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCalendar() *Calendar { _ = "STUB: not implemented"; return nil }
