package types

type HasPrivilegesUserProfileErrors struct {
	Count   int64                 `json:"count"`
	Details map[string]ErrorCause `json:"details"`
}

func (s *HasPrivilegesUserProfileErrors) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHasPrivilegesUserProfileErrors() *HasPrivilegesUserProfileErrors {
	_ = "STUB: not implemented"
	return nil
}
