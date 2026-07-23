package putcalendar

type Response struct {
	CalendarId string `json:"calendar_id"`

	Description *string `json:"description,omitempty"`

	JobIds []string `json:"job_ids"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
