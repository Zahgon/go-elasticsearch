package postcalendarevents

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Events []types.CalendarEvent `json:"events"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
