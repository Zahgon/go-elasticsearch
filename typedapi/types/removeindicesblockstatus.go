package types

type RemoveIndicesBlockStatus struct {
	Exception *ErrorCause `json:"exception,omitempty"`
	Name      string      `json:"name"`
	Unblocked *bool       `json:"unblocked,omitempty"`
}

func (s *RemoveIndicesBlockStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRemoveIndicesBlockStatus() *RemoveIndicesBlockStatus { _ = "STUB: not implemented"; return nil }
