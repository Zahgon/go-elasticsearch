package types

import (
	"encoding/json"
)

type ErrorCause struct {
	CausedBy *ErrorCause                `json:"caused_by,omitempty"`
	Metadata map[string]json.RawMessage `json:"-"`

	Reason    *string      `json:"reason,omitempty"`
	RootCause []ErrorCause `json:"root_cause,omitempty"`

	StackTrace *string      `json:"stack_trace,omitempty"`
	Suppressed []ErrorCause `json:"suppressed,omitempty"`

	Type string `json:"type"`
}

func (s *ErrorCause) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ErrorCause) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewErrorCause() *ErrorCause { _ = "STUB: not implemented"; return nil }
