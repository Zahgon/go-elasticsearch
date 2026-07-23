package poststarttrial

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type Response struct {
	Acknowledged    bool                     `json:"acknowledged"`
	ErrorMessage    *string                  `json:"error_message,omitempty"`
	TrialWasStarted bool                     `json:"trial_was_started"`
	Type            *licensetype.LicenseType `json:"type,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
