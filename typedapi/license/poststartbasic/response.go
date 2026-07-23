package poststartbasic

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type Response struct {
	Acknowledge     map[string][]string      `json:"acknowledge,omitempty"`
	Acknowledged    bool                     `json:"acknowledged"`
	BasicWasStarted bool                     `json:"basic_was_started"`
	ErrorMessage    *string                  `json:"error_message,omitempty"`
	Type            *licensetype.LicenseType `json:"type,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
