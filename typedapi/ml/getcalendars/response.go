package getcalendars

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Calendars []types.Calendar `json:"calendars"`
	Count     int64            `json:"count"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
