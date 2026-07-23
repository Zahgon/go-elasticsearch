package types

type ProjectEncryptionKeyDetails struct {
	ActiveKeyId        *string `json:"active_key_id,omitempty"`
	ActivePasswordId   string  `json:"active_password_id"`
	KeyCount           *int    `json:"key_count,omitempty"`
	MetadataPasswordId *string `json:"metadata_password_id,omitempty"`
	State              string  `json:"state"`
}

func (s *ProjectEncryptionKeyDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewProjectEncryptionKeyDetails() *ProjectEncryptionKeyDetails {
	_ = "STUB: not implemented"
	return nil
}
