package types

type GetUserProfileErrors struct {
	Count   int64                 `json:"count"`
	Details map[string]ErrorCause `json:"details"`
}

func (s *GetUserProfileErrors) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGetUserProfileErrors() *GetUserProfileErrors { _ = "STUB: not implemented"; return nil }
