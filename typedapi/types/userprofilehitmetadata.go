package types

type UserProfileHitMetadata struct {
	PrimaryTerm_ int64 `json:"_primary_term"`
	SeqNo_       int64 `json:"_seq_no"`
}

func (s *UserProfileHitMetadata) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUserProfileHitMetadata() *UserProfileHitMetadata { _ = "STUB: not implemented"; return nil }
